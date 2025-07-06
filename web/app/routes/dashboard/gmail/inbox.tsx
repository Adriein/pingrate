import React from "react";
import type {Route} from "~/.react-router/types/app/routes/landing/+types/landing";
import {COOKIE_HEADER, type SessionCookie, sessionCookie} from "@app/cookies-helper";
import {redirect} from "react-router";
import {getGmail, type GetGmailResponse, type Gmail, PingrateApiResponse} from "@app/shared/api/pingrate-api";

export async function loader({ request }: Route.LoaderArgs) {
    const { fromSession } = sessionCookie();

    const session: SessionCookie|null = await fromSession(request.headers.get(COOKIE_HEADER));

    if (!session) {
        return redirect("/signin");
    }

    const response: PingrateApiResponse<GetGmailResponse> = await getGmail(session.id);

    return {data: response.body?.data || []};
}

export default function GmailInbox({loaderData}: Route.ComponentProps) {
    const { data } = loaderData;
    return (
        <div>
            <p>Gmail Inbox</p>
            {data.map((item: Gmail, index: number) => (
                <div key={index}>
                    <p>{item.threadId}</p>
                    <p>{item.body}</p>
                </div>
            ))}
        </div>
    );
}