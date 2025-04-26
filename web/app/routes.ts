import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
    index("routes/landing/landing.tsx"),
    route("/signup", "routes/signup/signup.tsx")
] satisfies RouteConfig;
