import {PingrateApiResponse} from "@app/shared/api/PingrateApiResponse";

const PINGRATE_API_V1_URL: string = "http://localhost:4000/api/v1"

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
    return await post<SignupResponse>("/users/login", payload);
};


//SHARED
const post = async <T>(resource: string, payload: Record<string, any>): Promise<PingrateApiResponse<T>> => {
    try {
        const request: Request = new Request(`${PINGRATE_API_V1_URL}${resource}`, {
            method: "POST",
            body: JSON.stringify(payload),
        });

        const response: Response = await fetch(request);

        const body = await response.json();

        // response.ok only checks if the server responded with 2XX
        if (!response.ok) {
            return new PingrateApiResponse<T>(
                false,
                body,
            );
        }
        return new PingrateApiResponse<T>(true, body);
    } catch (error: unknown) {
        return new PingrateApiResponse<T>(
            false,
            undefined,
            error as Error
        );
    }

}