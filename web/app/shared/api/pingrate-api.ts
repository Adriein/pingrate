import {PingrateApiResponse} from "@app/shared/api/PingrateApiResponse";

export type SignupForm = { email: string, password: string};

const PINGRATE_API_V1_URL: string = "http://localhost:4000/api/v1"

export const signup = async (payload: SignupForm): Promise<PingrateApiResponse> => {
    return await post("/users", payload);
};

const post = async (resource: string, payload: Record<string, any>): Promise<PingrateApiResponse> => {
    try {
        const request: Request = new Request(`${PINGRATE_API_V1_URL}${resource}`, {
            method: "POST",
            body: JSON.stringify(payload),
        });

        const response: Response = await fetch(request);

        if (!response.ok) {
            return new PingrateApiResponse(
                true,
                await response.json(),
            );
        }

        return new PingrateApiResponse(true);
    } catch (error: unknown) {
        return new PingrateApiResponse(
            false,
            undefined,
            error as Error
        );
    }

}