import type {Route} from "~/.react-router/types/app/routes/landing/+types/landing";
import Integration from "@app/routes/dashboard/integration/integration";
import {Outlet, redirect} from "react-router";
import {COOKIE_HEADER, type SessionCookie, sessionCookie} from "@app/cookies-helper";
import {
    askGmailPermissions,
    type AskGmailPermissionsResponse,
    PingrateApiResponse,
} from "@app/shared/api/pingrate-api";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Pingrate | Dashboard" },
        { name: "description", content: "Signin page for Pingrate" },
    ];
}

export async function loader({ request }: Route.LoaderArgs) {
    const { fromSession } = sessionCookie();

    const session: SessionCookie|null = await fromSession(request.headers.get(COOKIE_HEADER));

    if (!session) {
        return redirect("/signin");
    }

    return {};
}

export async function action({request}: Route.ActionArgs){
    const { fromSession } = sessionCookie();

    const session: SessionCookie|null = await fromSession(request.headers.get(COOKIE_HEADER));

    if (!session) {
        return redirect("/signin");
    }

    const formData: FormData = await request.formData();

    switch (formData.get('integration')) {
        case 'gmail':
            const response: PingrateApiResponse<AskGmailPermissionsResponse> = await askGmailPermissions(session.id);

            if (!response.ok) {
                return redirect("/dashboard");
            }

            return redirect(response.body?.data!);
        default:
            break;
    }

    return {};
}

export default function Dashboard({loaderData}: Route.ComponentProps) {
    return (
        <div>
            <Integration/>
            <Outlet />
        </div>
    );
}



