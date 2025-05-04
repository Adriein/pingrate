import { PassThrough } from "node:stream";

import type { AppLoadContext, EntryContext } from "react-router";
import { createReadableStreamFromReadable } from "@react-router/node";
import { ServerRouter } from "react-router";
import { isbot } from "isbot";
import type { RenderToPipeableStreamOptions } from "react-dom/server";
import { renderToPipeableStream } from "react-dom/server";
import i18nextServer from "@app/language/i18n.server";
import {i18Config} from "@app/language/i18n.config";
import {createInstance} from "i18next";
import Backend from "i18next-fetch-backend";
import {I18nextProvider, initReactI18next} from "react-i18next";
import { resolve as fsresolve } from "node:path";

export const streamTimeout = 5_000;

export default function handleRequest(
  request: Request,
  responseStatusCode: number,
  responseHeaders: Headers,
  routerContext: EntryContext,
  loadContext: AppLoadContext
  // If you have middleware enabled:
  // loadContext: unstable_RouterContextProvider
) {
  return new Promise(async (resolve, reject) => {
    let shellRendered = false;
    let userAgent = request.headers.get("user-agent");

    // Ensure requests from bots and SPA Mode renders wait for all content to load before responding
    // https://react.dev/reference/react-dom/server/renderToPipeableStream#waiting-for-all-content-to-load-for-crawlers-and-static-generation
    let readyOption: keyof RenderToPipeableStreamOptions =
      (userAgent && isbot(userAgent)) || routerContext.isSpaMode
        ? "onAllReady"
        : "onShellReady";

    let instance = createInstance();

    let lng: string = await i18nextServer.getLocale(request);
    let ns: string[] = i18nextServer.getRouteNamespaces(routerContext);

    await instance
      .use(initReactI18next) // Tell our instance to use react-i18next
      .use(Backend)
      .init({
          ...i18Config, // spread the configuration
          lng, // The locale we detected above
          ns, // The namespaces the routes about to render wants to use
          backend: { loadPath: fsresolve("./public/locales/{{lng}}/{{ns}}.json") },
      });

    const { pipe, abort } = renderToPipeableStream(
        <I18nextProvider i18n={instance}>
            <ServerRouter context={routerContext} url={request.url} />
        </I18nextProvider>,
      {
        [readyOption]() {
          shellRendered = true;
          const body = new PassThrough();
          const stream = createReadableStreamFromReadable(body);

          responseHeaders.set("Content-Type", "text/html");

          resolve(
            new Response(stream, {
              headers: responseHeaders,
              status: responseStatusCode,
            })
          );

          pipe(body);
        },
        onShellError(error: unknown) {
          reject(error);
        },
        onError(error: unknown) {
          responseStatusCode = 500;
          // Log streaming rendering errors from inside the shell.  Don't log
          // errors encountered during initial shell rendering since they'll
          // reject and get logged in handleDocumentRequest.
          if (shellRendered) {
            console.error(error);
          }
        },
      }
    );

    // Abort the rendering stream after the `streamTimeout` so it has time to
    // flush down the rejected boundaries
    setTimeout(abort, streamTimeout + 1000);
  });
}
