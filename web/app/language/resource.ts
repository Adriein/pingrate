import commonSpanish from "@public/locales/es/common.json"
import commonEnglish from "@public/locales/en/common.json"

const languages = ["en", "es"] as const;
export const supportedLanguages = [...languages];
export type Language = (typeof languages)[number];

export type Resource = {
    common: typeof commonEnglish;
};

export const resources: Record<Language, Resource> = {
    en: {
        common: commonEnglish,
    },
    es: {
        common: commonSpanish,
    },
};

export const returnLanguageIfSupported = (
    lang?: string
): Language | undefined => {
    if (supportedLanguages.includes(lang as Language)) {
        return lang as Language;
    }
    return undefined;
};