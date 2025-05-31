import {Group} from "@mantine/core";
import {GmailButton} from "@app/routes/dashboard/integration/google";
import {useDisclosure} from "@mantine/hooks";

export default function Integration() {
    const [loading, { toggle }] = useDisclosure();
    return (
        <Group justify="center" p="md">
            <GmailButton
                loading={loading}
                loaderProps={{ type: 'dots' }}
            >
                Connect Gmail
            </GmailButton>
        </Group>
    );
}