import {type Cookie, createCookie} from "react-router";

export interface PingrateCookie extends Cookie {
    readonly value: string;
}

/**
 * Parses a cookie string into a PingrateCookie object
 * @param cookieString - The cookie string to parse (e.g. "$session=c5189e43-3c5e-4caa-9560-fe35ccd45d58; Path=/; Expires=Sun, 08 Jun 2025 15:47:46 GMT; HttpOnly; SameSite=Lax")
 * @returns A PingrateCookie object with the parsed values
 */
export async function parseCookie(cookieString: string): Promise<PingrateCookie> {
    const parts = cookieString.split(';').map(part => part.trim());

    const [nameValue,] = parts;
    const [name, value] = nameValue.split('=').map(part => part.trim());

    const remixCookie: Cookie = await createCookie(name).parse(cookieString);

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