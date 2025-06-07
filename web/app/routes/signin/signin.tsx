import React from "react";
import {type Cookie, data, Link, redirect, useFetcher} from "react-router";

import type { Route } from "./+types/signin";
import type { MantineTheme } from "@mantine/core";
import type { UseFormReturnType } from "@mantine/form";

import { useDisclosure } from "@mantine/hooks";
import { useMantineTheme } from "@mantine/core";

import {
    Anchor,
    Avatar,
    Button,
    Input,
    PasswordInput,
    Text,
    Title,
} from "@mantine/core";

import {
    hasLength,
    isEmail,
    isNotEmpty,
    useForm,
} from "@mantine/form";

import { IconAt, IconLock } from "@tabler/icons-react";

import PingrateLogo from "@app/shared/img/pingrate-logo.png";
import {PingrateApiResponse, signin, type SigninResponse, VALIDATION_ERROR} from "@app/shared/api/pingrate-api";
import classes from "./signin.module.css";
import PingrateError from "@app/shared/component/error/pingrate-error";
import {sessionCookie} from "@app/cookies.server";
import {type SigninTranslations, translate} from "@app/locale.server";
import {ES} from "@app/shared/constants";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Pingrate" },
        { name: "description", content: "Signin page for Pingrate" },
    ];
}

export async function loader({ context, request }: Route.LoaderArgs) {
    const cookie: Cookie|null = await sessionCookie.parse(request.headers.get('set-cookie'));

    console.log(cookie);

    const translations: SigninTranslations = translate(ES, "signin");

    return {
        lang: {...translations}
    }
}

export async function action({request}: Route.ActionArgs){
    const formData: FormData = await request.formData();

    const response: PingrateApiResponse<SigninResponse> = await signin({
        email: formData.get('email') as string,
        password: formData.get('password') as string
    });

    if (!response.ok) {
        if (response.data?.error === VALIDATION_ERROR) {
            return data({ error: response.data.data }, { status: 400 });
        }

        return data({ error: "Something went wrong" }, { status: 500 });
    }

    console.log(response)

    return {};

    /*return redirect("/dashboard", {
        headers: {
            "Set-Cookie": await sessionCookie.serialize({ sessionId: response.cookieSession }),
        },
    });*/
}

export default function Signin({loaderData}: Route.ComponentProps) {
    const { lang } = loaderData;
    const fetcher = useFetcher();
    const theme: MantineTheme = useMantineTheme();
    const [visible, { toggle }] = useDisclosure(false);

    const error = fetcher.data?.error;


    const form: UseFormReturnType<any> = useForm({
        mode: 'uncontrolled',
        initialValues: {
            email: '',
            password: '',
            termsOfService: true,
        },
        validate: {
            email: (value: string) => {
                const validations = [
                    isNotEmpty('Must add an email'),
                    isEmail('Invalid email')
                ];

                const results: React.ReactNode[] = validations
                    .map(validation => validation(value))
                    .filter((result: React.ReactNode | null) => !!result);

                return results[0];
            },
            password: (value: string) => {
                const validations = [
                    isNotEmpty('Must add a password'),
                    hasLength({min: 8, max: 20}, 'Password must be at least 8 char long and max 20 char long'),
                ];

                const results: React.ReactNode[] = validations
                    .map(validation => validation(value))
                    .filter((result: React.ReactNode | null) => !!result);

                return results[0];
            },
        }
    });

    const handleSubmit = async (values: typeof form.values): Promise<void> => {
        await fetcher.submit(values, {method: 'POST'})
    };

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
                    {lang.title}
                </Title>
            </div>
            <div className={classes.containerForm}>
                <fetcher.Form onSubmit={form.onSubmit(handleSubmit)} className={classes.form}>
                    <Input.Wrapper
                        label="Email"
                        withAsterisk
                        key={form.key('email')}
                        error={form.getInputProps('email').error ?? error?.email}
                        styles={{
                            label: {
                                color: theme.colors.pingrateSecondary[10]
                            }
                        }}
                    >
                        <Input
                            placeholder="example@gmail.com"
                            leftSection={<IconAt size={16} />}
                            classNames={classes}
                            error={form.getInputProps('email').error ?? error?.email}
                            onChange={form.getInputProps('email').onChange}
                            onBlur={form.getInputProps('email').onBlur}
                            onFocus={form.getInputProps('email').onFocus}
                            defaultValue={form.getInputProps('email').defaultValue}
                        />
                    </Input.Wrapper>
                    <Input.Wrapper
                        label="Password"
                        withAsterisk
                        key={form.key('password')}
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
                            classNames={classes}
                            {...form.getInputProps('password')}
                        />
                    </Input.Wrapper>
                    <Button
                        fullWidth
                        variant="filled"
                        type="submit"
                        loading={fetcher.state !== "idle"}
                        vars={(theme: MantineTheme) => ({
                            root: {
                                '--button-bg': theme.colors.pingratePrimary[6],
                                '--button-hover': theme.colors.pingratePrimary[7],
                            },
                        })}
                    >
                        {lang.button}
                    </Button>
                </fetcher.Form>
                <div className={classes.formLink}>
                    <Text size="sm" c="dimmed">{lang.signupShortcut}</Text>
                    <Anchor
                        component={Link}
                        to="/signup"
                        size="sm"
                        styles={{
                            root: {
                                color: theme.colors.pingratePrimary[10]
                            }
                        }}
                    >
                        {lang.signupShortcutLink}
                    </Anchor>
                </div>
            </div>
            {error && typeof error === 'string' && <PingrateError message={error}/>}
        </div>
    );
}