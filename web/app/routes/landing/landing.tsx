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
import {type LandingTranslations, translate} from "@app/locale.server";
import {ES} from "@app/shared/constants";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Pingrate" },
        { name: "description", content: "Landing page for Pingrate" },
    ];
}
export async function loader({ context }: Route.LoaderArgs) {
    const translations: LandingTranslations = translate(ES, "landing");
    return { lang: {...translations} }
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
                        <Link to="/signin">
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
                                {lang.loginButton}
                            </Button>
                        </Link>
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
                                {lang.signupButton}
                            </Button>
                        </Link>
                    </div>
                </div>
            </nav>
        </main>
    );
}
