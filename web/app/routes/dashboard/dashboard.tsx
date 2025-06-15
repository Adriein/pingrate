import type {Route} from "~/.react-router/types/app/routes/landing/+types/landing";
import {data, Outlet, redirect} from "react-router";
import {COOKIE_HEADER, type SessionCookie, sessionCookie} from "@app/cookies-helper";
import {
    askGmailPermissions,
    type AskGmailPermissionsResponse,
    PingrateApiResponse,
} from "@app/shared/api/pingrate-api";
import React from "react";
import {ActionIcon, Group, Tooltip, Text, Avatar, Code, Title, type MantineTheme, useMantineTheme} from "@mantine/core";
import {IconChevronRight, IconPlus} from "@tabler/icons-react";
import classes from './dashboard.module.css';
import PingrateLogo from "@app/shared/img/pingrate-logo.png";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Pingrate | Dashboard" },
        { name: "description", content: "Signin page for Pingrate" },
    ];
}

export async function loader({ request }: Route.LoaderArgs) {
    const { fromSession } = sessionCookie();

    const session: SessionCookie|null = await fromSession(request.headers.get(COOKIE_HEADER));

    if (!session) {
        return redirect("/signin");
    }

    return {};
}

export async function action({request}: Route.ActionArgs){
    const { fromSession } = sessionCookie();

    const session: SessionCookie|null = await fromSession(request.headers.get(COOKIE_HEADER));

    if (!session) {
        return redirect("/signin");
    }

    const formData: FormData = await request.formData();

    switch (formData.get('integration')) {
        case 'gmail':
            const response: PingrateApiResponse<AskGmailPermissionsResponse> = await askGmailPermissions(session.id);

            if (!response.ok) {
                return data({ error: "Something went wrong" }, { status: 500 });
            }

            return redirect(response.body?.data!);
        default:
            break;
    }

    return {};
}

export default function Dashboard({loaderData}: Route.ComponentProps) {
    const theme: MantineTheme = useMantineTheme();

    return (
        <div className={classes.main}>
            <nav className={classes.navbar}>
                <div className={classes.header}>
                    <Group justify="space-between">
                        <Avatar src={PingrateLogo} size={36} alt="Pingrate logo"/>
                        <Title
                            order={2}
                            styles={{
                                root: {
                                    color: theme.colors.pingrateSecondary[10]
                                }
                            }}
                        >
                            Pingrate
                        </Title>
                        <Code fw={700}>v0.1.0</Code>
                    </Group>
                </div>

                <div className={classes.section}>

                </div>

                <div className={classes.section}>
                    <Group className={classes.collectionsHeader} justify="space-between">
                        <Text size="xs" fw={500} c="dimmed">
                            Collections
                        </Text>
                        <Tooltip label="Create collection" withArrow position="right">
                            <ActionIcon variant="default" size={18}>
                                <IconPlus size={12} stroke={1.5} />
                            </ActionIcon>
                        </Tooltip>
                    </Group>
                </div>

                <div className={classes.section}>
                    <div className={classes.user}>
                        <Group>
                            <Avatar
                                src="https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/avatars/avatar-8.png"
                                radius="xl"
                            />

                            <div style={{ flex: 1 }}>
                                <Text size="sm" fw={500}>
                                    Harriette Spoonlicker
                                </Text>

                                <Text c="dimmed" size="xs">
                                    hspoonlicker@outlook.com
                                </Text>
                            </div>

                            <IconChevronRight size={14} stroke={1.5} />
                        </Group>
                    </div>
                </div>
            </nav>
            <div className={classes.content}>
                <h1>Dashboard</h1>
                <Outlet />
            </div>
        </div>
    );
}



