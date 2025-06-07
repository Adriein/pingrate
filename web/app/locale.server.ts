import * as spanish from "@app/locales/es";
import * as english from "@app/locales/en";
import {ES} from "@app/shared/constants";

export interface LandingTranslations {
    signupButton: string;
    loginButton: string;
}

export interface SignupTranslations {
    title: string;
    button: string;
    termsAndPrivacy: string;
    termsAndPrivacyLink: string;
    loginShortcut: string;
    loginShortcutLink: string;
}

export interface SigninTranslations {
    title: string;
    button: string;
    signupShortcut: string;
    signupShortcutLink: string;
}

export type PageTranslations = {
    landing: LandingTranslations;
    signup: SignupTranslations;
    signin: SigninTranslations;
}

// Define a generic type that returns the correct translation interface based on the page parameter
type TranslationForPage<P extends keyof PageTranslations> = PageTranslations[P];

export const translate: <P extends keyof PageTranslations>(lng: string, page: P) => TranslationForPage<P> =
    <P extends keyof PageTranslations>(lng: string, page: P): TranslationForPage<P> => {
        if (lng === ES) {
            return spanish.default[page] as TranslationForPage<P>;
        }

        return english.default[page] as TranslationForPage<P>;
}
