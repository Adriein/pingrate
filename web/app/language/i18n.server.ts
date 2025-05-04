import { resolve } from "node:path";
import { RemixI18Next } from "remix-i18next/server";
import {i18Config} from "@app/language/i18n.config";
import Backend from "i18next-fetch-backend";

let i18nextServer: RemixI18Next = new RemixI18Next({
    detection: {
        supportedLanguages: i18Config.supportedLngs,
        fallbackLanguage: i18Config.fallbackLng,
    },
    // This is the configuration for i18next used
    // when translating messages server-side only
    i18next: {
        ...i18Config,
        backend: {
            loadPath: resolve("./public/locales/{{lng}}/{{ns}}.json"),
        },
    },
    // The i18next plugins you want RemixI18next to use for `i18n.getFixedT` inside loaders and actions.
    // E.g. The Backend plugin for loading translations from the file system
    // Tip: You could pass `resources` to the `i18next` configuration and avoid a backend here
    plugins: [Backend],
});

export default i18nextServer;