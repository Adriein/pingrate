import {Group} from "@mantine/core";
import {GmailButton} from "@app/routes/dashboard/integration/google";
import {useFetcher} from "react-router";
import PingrateError from "@app/shared/component/error/pingrate-error";
import React from "react";

export default function Integration() {
    const fetcher = useFetcher();

    const error = fetcher.data?.error;

    const handleConnectGmail = async () => {
        await fetcher.submit(
            { integration: "gmail" },
            { method: "POST" }
        );
    };

    return (
        <div>
            <Group justify="center" p="md">
                <GmailButton
                    loaderProps={{ type: 'dots' }}
                    loading={fetcher.state !== "idle"}
                    onClick={handleConnectGmail}
                >
                    Connect Gmail
                </GmailButton>
            </Group>
            {error && typeof error === 'string' && <PingrateError message={error}/>}
        </div>
    );
}
