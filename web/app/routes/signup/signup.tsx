import type { Route } from "./+types/signup";
import {useDisclosure} from "@mantine/hooks";
import {
    Anchor,
    Avatar,
    Button,
    Checkbox,
    Input,
    type MantineTheme,
    PasswordInput,
    Text,
    Title,
    useMantineTheme
} from "@mantine/core";
import { IconAt, IconLock } from '@tabler/icons-react';
import PingrateLogo from "@app/shared/img/pingrate-logo.png";
import classes from "./signup.module.css";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Pingrate" },
        { name: "description", content: "Signup page for Pingrate" },
    ];
}

export default function Signup() {
    const theme: MantineTheme = useMantineTheme();
    const [visible, { toggle }] = useDisclosure(false);
    return (
        <div className={classes.container}>
            <div className={classes.containerHeader}>
                <Avatar className={classes.containerHeaderLogo} src={PingrateLogo} size={36} alt="Pingrate logo"/>
                <Title
                    className={classes.containerHeaderTitle}
                    order={2}
                    styles={{
                        root: {
                            color: theme.colors.pingrateSecondary[10]
                        }
                    }}
                >
                    Join Pingrate
                </Title>
            </div>
            <div className={classes.containerForm}>
                <form className={classes.form} action="#" method="POST">
                    <Input.Wrapper
                        label="Email"
                        error="Input error"
                        styles={{
                            label: {
                                color: theme.colors.pingrateSecondary[10]
                            }
                        }}
                    >
                        <Input
                            placeholder="example@gmail.com"
                            leftSection={<IconAt size={16} />}
                        />
                    </Input.Wrapper>
                    <Input.Wrapper
                        label="Password"
                        error="Input error"
                        styles={{
                            label: {
                                color: theme.colors.pingrateSecondary[10]
                            }
                        }}
                    >
                        <PasswordInput
                            leftSection={<IconLock size={16} />}
                            visible={visible}
                            onVisibilityChange={toggle}
                        />
                    </Input.Wrapper>
                    <Checkbox
                        defaultChecked
                        label={
                            <>
                                I agree to the{' '}
                                <Anchor
                                    href="https://mantine.dev"
                                    target="_blank"
                                    inherit
                                    styles={{
                                        root: {
                                            color: theme.colors.pingratePrimary[10]
                                        }
                                    }}
                                >
                                    Terms & Privacy
                                </Anchor>
                            </>
                        }
                        vars={(theme: MantineTheme) => ({
                            root: {
                                '--checkbox-color': theme.colors.pingratePrimary[6],
                            },
                        })}
                        styles={{
                            label: {
                                color: theme.colors.pingrateSecondary[10]
                            }
                        }}
                    />
                    <Button
                        fullWidth
                        variant="filled"
                        vars={(theme: MantineTheme) => ({
                            root: {
                                '--button-bg': theme.colors.pingratePrimary[6],
                                '--button-hover': theme.colors.pingratePrimary[7],
                            },
                        })}
                    >
                        Sign up
                    </Button>
                </form>
                <div className={classes.formLink}>
                    <Text size="sm" c="dimmed">Do you have an account?</Text>
                    <Anchor
                        size="sm"
                        styles={{
                            root: {
                                color: theme.colors.pingratePrimary[10]
                            }
                        }}
                    >
                        Sign in
                    </Anchor>
                </div>
            </div>
        </div>
    );
}