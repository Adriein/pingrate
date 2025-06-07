import {type Cookie, createCookie} from "react-router";

export type PingrateCookie = {
    name: string;
    value: string;
    expires: Date;
    path: string;
    domain: string;
    secure: boolean;
    httpOnly: boolean;
    sameSite: "lax" | "strict" | "none";
}

/**
 * Parses a cookie string into a PingrateCookie object
 * @param cookieString - The cookie string to parse (e.g. "$session=c5189e43-3c5e-4caa-9560-fe35ccd45d58; Path=/; Expires=Sun, 08 Jun 2025 15:47:46 GMT; HttpOnly; SameSite=Lax")
 * @returns A PingrateCookie object with the parsed values
 */
export function parseCookie(cookieString: string): PingrateCookie {
    const parts = cookieString.split(';').map(part => part.trim());

    const [nameValue, ...attributes] = parts;
    const [name, value] = nameValue.split('=').map(part => part.trim());

    const cookie: PingrateCookie = {
        name,
        value,
        expires: new Date(Date.now() + 604_800 * 1000), // Default: one week from now
        path: '/',
        domain: '',
        secure: false,
        httpOnly: false,
        sameSite: 'lax'
    };

    // Parse the attributes
    for (const attr of attributes) {
        if (attr.toLowerCase().startsWith('path=')) {
            cookie.path = attr.substring(5).trim();
        }

        if (attr.toLowerCase().startsWith('expires=')) {
            cookie.expires = new Date(attr.substring(8).trim());
        }

        if (attr.toLowerCase().startsWith('domain=')) {
            cookie.domain = attr.substring(7).trim();
        }

        if (attr.toLowerCase() === 'secure') {
            cookie.secure = true;
        }

        if (attr.toLowerCase() === 'httponly') {
            cookie.httpOnly = true;
        }

        if (attr.toLowerCase().startsWith('samesite=')) {
            const sameSiteValue = attr.substring(9).trim().toLowerCase();
            if (sameSiteValue === 'lax' || sameSiteValue === 'strict' || sameSiteValue === 'none') {
                cookie.sameSite = sameSiteValue as "lax" | "strict" | "none";
            }
        }
    }

    return cookie;
}

export const sessionCookie = (pingrateCookie: PingrateCookie): Cookie => {
    return createCookie(pingrateCookie.name, {
        httpOnly: pingrateCookie.httpOnly,
        maxAge: 604_800, // one week
    });
}
