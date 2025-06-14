import {Group} from "@mantine/core";
import {GmailButton} from "@app/routes/dashboard/integration/google";
import {useFetcher} from "react-router";

export default function Integration() {
    const fetcher = useFetcher();

    const handleConnectGmail = async () => {
        await fetcher.submit(
            { integration: "gmail" },
            { method: "POST" }
        );
    };

    return (
        <Group justify="center" p="md">
            <GmailButton
                loaderProps={{ type: 'dots' }}
                loading={fetcher.state !== "idle"}
                onClick={handleConnectGmail}
            >
                Connect Gmail
            </GmailButton>
        </Group>
    );
}
