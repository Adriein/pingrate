import type {Route} from "~/.react-router/types/app/routes/landing/+types/landing";
import Integration from "@app/routes/dashboard/integration/integration";
import {Outlet, type Session} from "react-router";
import {askForGmailPermissions} from "@app/shared/api/pingrate-api";

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
    const formData: FormData = await request.formData();

    switch (formData.get('integration')) {
        case 'gmail':
            await askForGmailPermissions();
            break;
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



