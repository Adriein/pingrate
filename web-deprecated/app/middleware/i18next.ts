import { initReactI18next } from "react-i18next";
import {type Cookie, createCookie} from "react-router";
import { unstable_createI18nextMiddleware } from "remix-i18next/middleware";
import en from "@app/locales/en";
import es from "@app/locales/es";

export const localeCookie: Cookie = createCookie("lng", {
    path: "/",
    sameSite: "lax",
    secure: process.env.NODE_ENV === "production",
    httpOnly: true,
});

export const [i18nextMiddleware, getLocale, getInstance] =
    unstable_createI18nextMiddleware({
        detection: {
            supportedLanguages: ["es", "en"],
            fallbackLanguage: "en",
            cookie: localeCookie,
        },
        i18next: {
            resources: { en: { translation: en }, es: { translation: es } },
        },
        plugins: [initReactI18next],
    });

declare module "i18next" {
    interface CustomTypeOptions {
        resources: {
            translation: typeof en;
        };
    }
}