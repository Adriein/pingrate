import type {Route} from "~/.react-router/types/app/routes/landing/+types/landing";
import Integration from "@app/routes/dashboard/integration/integration";
import {Outlet, redirect} from "react-router";
import {type SessionCookie, sessionCookie} from "@app/cookies-helper";
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
    return {};
}

export async function action({request}: Route.ActionArgs){
    const { fromSession } = sessionCookie();
    const cookieSession: SessionCookie|null = await fromSession(request.headers.get('cookie'));

    if (!cookieSession) {
        //TODO: return unauthorized error and redirect
        return {};
    }

    const formData: FormData = await request.formData();

    switch (formData.get('integration')) {
        case 'gmail':
            const response: PingrateApiResponse<AskGmailPermissionsResponse> = await askGmailPermissions(cookieSession.id);

            return redirect(response.data?.data!);
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



