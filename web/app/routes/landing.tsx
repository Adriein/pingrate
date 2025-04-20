import type { Route } from "./+types/home";
import { Landing } from "@app/landing/landing";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Pingrate" },
    { name: "description", content: "Landing page for Pingrate" },
  ];
}

export default function Home() {
  return <Landing />;
}
