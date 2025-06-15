import {type Cookie, type CookieSerializeOptions, createCookie} from "react-router";

export interface PingrateCookie extends Cookie {
    readonly value: string;
}

export const COOKIE_HEADER = "cookie";

/**
 * Parses a cookie string into a PingrateCookie object
 * @param cookieString - The cookie string to parse (e.g. "$session=c5189e43-3c5e-4caa-9560-fe35ccd45d58; Path=/; Expires=Sun, 08 Jun 2025 15:47:46 GMT; HttpOnly; SameSite=Lax")
 * @returns A PingrateCookie object with the parsed values
 */
export async function parseCookie(cookieString: string): Promise<PingrateCookie> {
    const parts = cookieString.split(';').map(part => part.trim());

    const [nameValue, ...attributes] = parts;
    const [name, value] = nameValue.split('=').map(part => part.trim());

    const options: CookieSerializeOptions = {
        expires: undefined,
        path: undefined,
        domain: undefined,
        secure: undefined,
        httpOnly: undefined,
        sameSite: undefined
    };

    // Parse the attributes
    for (const attr of attributes) {
        if (attr.toLowerCase().startsWith('path=')) {
            options.path = attr.substring(5).trim();
        }

        if (attr.toLowerCase().startsWith('expires=')) {
            options.expires = new Date(attr.substring(8).trim());
        }

        if (attr.toLowerCase().startsWith('domain=')) {
            options.domain = attr.substring(7).trim();
        }

        if (attr.toLowerCase() === 'secure') {
            options.secure = true;
        }

        if (attr.toLowerCase() === 'httponly') {
            options.httpOnly = true;
        }

        if (attr.toLowerCase().startsWith('samesite=')) {
            const sameSiteValue = attr.substring(9).trim().toLowerCase();
            if (sameSiteValue === 'lax' || sameSiteValue === 'strict' || sameSiteValue === 'none') {
                options.sameSite = sameSiteValue as "lax" | "strict" | "none";
            }
        }
    }

    const remixCookie: Cookie = createCookie(name, options);

    return {value, ...remixCookie};
}

export interface SessionCookie extends Cookie {
    id: string;
}

export const sessionCookie = () => {
    const $name: string = "$session";

    const fromCookie: (cookie: PingrateCookie) => Promise<string> = async (cookie: PingrateCookie): Promise<string> => {
        const { name, value, ...remixCookieFormat } = cookie;

        const remixCookie: Cookie = createCookie(name, remixCookieFormat);

        return await remixCookie.serialize({id: value});
    }

    const fromSession: (session: string|null) => Promise<SessionCookie|null> = async (session: string|null): Promise<SessionCookie|null> => {
        return await createCookie($name).parse(session);
    }

    return { fromCookie, fromSession };
}