import {
    Avatar,
    Button,
    type MantineTheme, Title,
    useMantineTheme
} from "@mantine/core";
import PingrateLogo from "../shared/img/pingrate-logo.png";
import classes from "./landing.module.css";

export function Landing() {
    const theme: MantineTheme = useMantineTheme();
    return (
        <main className={classes.mainContainer}>
            <nav className={classes.header}>
                <div className={classes.headerContent}>
                    <div className={classes.headerTitle}>
                        <Avatar src={PingrateLogo} alt="Pingrate logo"/>
                        <Title
                            order={2}
                            styles={{
                                root: {
                                    color: theme.colors.pingrateSecondary[10]
                                }
                            }}
                        >Pingrate</Title>
                    </div>

                    <Button
                        variant="default"
                        vars={(theme: MantineTheme) => ({
                            root: {
                                '--button-bg': theme.colors.pingrateBackground[10],
                                '--button-hover': theme.colors.pingrateSecondary[0],
                                '--button-color': theme.colors.pingrateSecondary[10],
                            },
                        })}
                    >Login</Button>
                </div>
            </nav>
        </main>
    );
}
