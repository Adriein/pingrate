import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
    index("routes/landing/landing.tsx"),
    route("/signup", "routes/signup/signup.tsx"),
    route("/signin", "routes/signin/signin.tsx"),
    route("/dashboard", "routes/dashboard/dashboard.tsx", [
        route("/dashboard/integration", "routes/dashboard/integration/integration.tsx"),
    ]),
] satisfies RouteConfig;
