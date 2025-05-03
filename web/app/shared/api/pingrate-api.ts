import {PingrateApiResponse} from "@app/shared/api/PingrateApiResponse";

export type SignupForm = { id: string, email: string, password: string};
export type SignupResponse = { ok: boolean, data: { } | undefined, error: string | undefined }

const PINGRATE_API_V1_URL: string = "http://localhost:4000/api/v1"

export const signup = async (payload: SignupForm): Promise<PingrateApiResponse<SignupResponse>> => {
    return await post<SignupResponse>("/users", payload);
};

const post = async <T>(resource: string, payload: Record<string, any>): Promise<PingrateApiResponse<T>> => {
    try {
        const request: Request = new Request(`${PINGRATE_API_V1_URL}${resource}`, {
            method: "POST",
            body: JSON.stringify(payload),
        });

        const response: Response = await fetch(request);

        console.log(await response.json())
        console.log(response)

        // response.ok only checks if the server responded with 2XX
        if (!response.ok) {
            return new PingrateApiResponse<T>(
                true,
                await response.json(),
            );
        }

        return new PingrateApiResponse<T>(true);
    } catch (error: unknown) {
        return new PingrateApiResponse<T>(
            false,
            undefined,
            error as Error
        );
    }

}