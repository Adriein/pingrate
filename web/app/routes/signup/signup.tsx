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
                    <Input.Wrapper label="Email" error="Input error">
                        <Input
                            placeholder="example@gmail.com"
                            leftSection={<IconAt size={16} />}
                        />
                    </Input.Wrapper>
                    <Input.Wrapper label="Password" error="Input error">
                        <PasswordInput
                            leftSection={<IconLock size={16} />}
                            visible={visible}
                            onVisibilityChange={toggle}
                        />
                    </Input.Wrapper>
                    <Checkbox
                        defaultChecked
                        label="I agree to the Terms & Privacy"
                        vars={(theme: MantineTheme) => ({
                            root: {
                                '--checkbox-color': theme.colors.pingrateAccent[5],
                            },
                        })}
                    />
                    <Button
                        fullWidth
                        variant="filled"
                        vars={(theme: MantineTheme) => ({
                            root: {
                                '--button-bg': theme.colors.pingrateAccent[5],
                                '--button-hover': theme.colors.pingrateAccent[7],
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
                                color: theme.colors.pingrateAccent[5]
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