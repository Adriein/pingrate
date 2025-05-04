import {
    Avatar,
    Button,
    type MantineTheme, Title,
    useMantineTheme
} from "@mantine/core";
import PingrateLogo from "../../shared/img/pingrate-logo.png";
import classes from "./landing.module.css";

import type { Route } from "./+types/landing";
import {Link} from "react-router";
import i18nextServer from "@app/language/i18n.server"
import type {TFunction} from "i18next";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Pingrate" },
        { name: "description", content: "Landing page for Pingrate" },
    ];
}
export async function loader({ request }: Route.LoaderArgs) {
    const t: TFunction = await i18nextServer.getFixedT("es", "landing");
    const signup = t("signup");

    console.log(signup)

    return { lang: { signup } }
}


export default function Home({loaderData}: Route.ComponentProps) {
    const { lang } = loaderData;

    const theme: MantineTheme = useMantineTheme();
    return (
        <main className={classes.mainContainer}>
            <nav className={classes.header}>
                <div className={classes.headerContent}>
                    <div className={classes.headerTitle}>
                        <Avatar src={PingrateLogo} size={36} alt="Pingrate logo"/>
                        <Title
                            order={2}
                            styles={{
                                root: {
                                    color: theme.colors.pingrateSecondary[10]
                                }
                            }}
                        >Pingrate</Title>
                    </div>
                    <div className={classes.headerButtons}>
                        <Button
                            variant="default"
                            vars={(theme: MantineTheme) => ({
                                root: {
                                    '--button-bg': theme.colors.pingrateBackground[10],
                                    '--button-hover': theme.colors.pingrateSecondary[0],
                                    '--button-color': theme.colors.pingrateSecondary[10],
                                },
                            })}
                        >
                            Sign In
                        </Button>
                        <Link to="/signup">
                            <Button
                                variant="filled"
                                radius="xl"
                                vars={(theme: MantineTheme) => ({
                                    root: {
                                        '--button-bg': theme.colors.pingrateAccent[5],
                                        '--button-hover': theme.colors.pingrateAccent[7],
                                    },
                                })}
                            >
                                {lang.signup}
                            </Button>
                        </Link>
                    </div>
                </div>
            </nav>
        </main>
    );
}
