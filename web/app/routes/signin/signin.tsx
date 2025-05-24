import React from "react";
import {data, Link, redirect, useFetcher} from "react-router";
import { v4 as uuidv4 } from 'uuid';

import type { Route } from "./+types/signin";
import type { PingrateApiResponse } from "@app/shared/api/PingrateApiResponse";
import type { MantineTheme } from "@mantine/core";
import type { UseFormReturnType } from "@mantine/form";

import { useDisclosure } from "@mantine/hooks";
import { useMantineTheme } from "@mantine/core";

import {
    Anchor,
    Avatar,
    Button,
    Checkbox,
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
import {signup, type SignupResponse, VALIDATION_ERROR} from "@app/shared/api/pingrate-api";
import classes from "./signin.module.css";
import PingrateError from "@app/shared/component/error/pingrate-error";
import type {i18n} from "i18next";
import {getInstance} from "@app/middleware/i18next";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Pingrate" },
        { name: "description", content: "Signup page for Pingrate" },
    ];
}

export async function loader({ context }: Route.LoaderArgs) {
    const i18next: i18n = getInstance(context);

    const title: string = i18next.t("signin.title");
    const signupButton: string = i18next.t("signin.button");
    const signupShortcut: string = i18next.t("signin.signupShortcut");
    const signupShortcutLink: string = i18next.t("signin.signupShortcutLink");

    return {
        lang: {
            title,
            signupButton,
            signupShortcut,
            signupShortcutLink,
        }
    }
}

export async function action({request}: Route.ActionArgs){
    const formData: FormData = await request.formData();

    const response: PingrateApiResponse<SignupResponse> = await signup({
        id: uuidv4(),
        email: formData.get('email') as string,
        password: formData.get('password') as string
    });

    if (!response.ok) {
        if (response.data?.error === VALIDATION_ERROR) {
            return data({ error: response.data.data }, { status: 400 });
        }

        return data({ error: "Something went wrong" }, { status: 500 });
    }

    return redirect("/dashboard");
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
                        {lang.signupButton}
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