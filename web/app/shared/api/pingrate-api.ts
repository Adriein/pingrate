import {parseCookie, type PingrateCookie} from "@app/cookies-helper";

const PINGRATE_API_V1_URL: string = "http://localhost:4000/api/v1"

export class PingrateApiResponse<T = undefined> {
    public static error<T>(body?: T, error?: Error) {
        return new PingrateApiResponse<T>(false, body, [], error);
    }
    public constructor(
        private _ok: boolean,
        private _body?: T,
        private _cookies: PingrateCookie[] = [],
        private _error?: Error,
    ) {}

    public get ok(): boolean {
        return this._ok;
    }

    public get body(): T | undefined {
        return this._body;
    }

    public get cookies(): PingrateCookie[] {
        return this._cookies;
    }

    public get error(): Error | undefined {
        return this._error;
    }
}

//SIGNUP
export type SignupForm = { id: string, email: string, password: string};
export type SignupResponse = { ok: boolean, data: Record<string, string> | undefined, error: string | undefined }

export const VALIDATION_ERROR: string = "VALIDATION_ERROR";

export const signup = async (payload: SignupForm): Promise<PingrateApiResponse<SignupResponse>> => {
    return await post<SignupResponse>("/users", payload);
};

//SIGNIN
export type SigninForm = { email: string, password: string};
export type SigninResponse = { ok: boolean, data: Record<string, string> | undefined, error: string | undefined }

export const signin = async (payload: SigninForm): Promise<PingrateApiResponse<SigninResponse>> => {
    return await post<SignupResponse>("/auth", payload);
};

//GMAIL INTEGRATION
export type AskGmailPermissionsResponse = { ok: boolean, data: string, error: string | undefined }

export const askGmailPermissions = async (sessionId: string): Promise<PingrateApiResponse<AskGmailPermissionsResponse>> => {
    return await get("/integrations/gmail/oauth", { 'Cookie': `$session=${sessionId}` });
};

export type GetGmailResponse = { ok: boolean, data: Gmail[], error: string | undefined }
export type Gmail = {id: string, threadId: string, body: string};

export const getGmail = async (sessionId: string): Promise<PingrateApiResponse<GetGmailResponse>> => {
    return await get("/integrations/gmail", { 'Cookie': `$session=${sessionId}` });
}



//SHARED
const get = async <T>(resource: string, headers: Record<string, string>): Promise<PingrateApiResponse<T>> => {
    try {
        const request: Request = new Request(`${PINGRATE_API_V1_URL}${resource}`, {
            method: "GET",
            credentials: 'include',
            headers: headers
        });

        const response: Response = await fetch(request);

        const body: T = await response.json();

        // response.ok only checks if the server responded with 2XX
        if (!response.ok) {
            return PingrateApiResponse.error<T>(body);
        }

        const cookieHeader: string|null = response.headers.get("set-cookie");

        if (cookieHeader) {
            const cookies: PingrateCookie[] = [await parseCookie(cookieHeader)];

            return new PingrateApiResponse<T>(true, body, cookies);
        }

        return new PingrateApiResponse<T>(true, body);
    } catch (error: unknown) {
        return PingrateApiResponse.error<T>(undefined, error as Error);
    }
}

const post = async <T>(resource: string, payload: Record<string, any>): Promise<PingrateApiResponse<T>> => {
    try {
        const request: Request = new Request(`${PINGRATE_API_V1_URL}${resource}`, {
            headers: {
                "Content-Application": "application/json"
            },
            method: "POST",
            body: JSON.stringify(payload),
        });

        const response: Response = await fetch(request);

        const body = await response.json();

        // response.ok only checks if the server responded with 2XX
        if (!response.ok) {
            return PingrateApiResponse.error<T>(body);
        }

        const cookieHeader: string|null = response.headers.get("set-cookie");

        if (cookieHeader) {
            const cookies: PingrateCookie[] = [await parseCookie(cookieHeader)];

            return new PingrateApiResponse<T>(true, body, cookies);
        }

        return new PingrateApiResponse<T>(true, body);
    } catch (error: unknown) {
        return PingrateApiResponse.error<T>(undefined, error as Error);
    }
}