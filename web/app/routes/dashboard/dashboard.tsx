import type {Route} from "~/.react-router/types/app/routes/landing/+types/landing";
import Integration from "@app/routes/dashboard/integration/integration";
import {Outlet} from "react-router";
import {whoAmI} from "@app/shared/api/pingrate-api";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Pingrate | Dashboard" },
        { name: "description", content: "Signin page for Pingrate" },
    ];
}

export async function loader({ params }: Route.LoaderArgs) {
    await whoAmI();
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



