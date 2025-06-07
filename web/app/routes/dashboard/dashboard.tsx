import type {Route} from "~/.react-router/types/app/routes/landing/+types/landing";
import Integration from "@app/routes/dashboard/integration/integration";
import {Outlet, type Session} from "react-router";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Pingrate | Dashboard" },
        { name: "description", content: "Signin page for Pingrate" },
    ];
}

export async function loader({ request }: Route.LoaderArgs) {
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



