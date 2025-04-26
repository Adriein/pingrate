import type { Route } from "./+types/signup";
import {useDisclosure} from "@mantine/hooks";
import {Button, Checkbox, Input, PasswordInput, Stack} from "@mantine/core";
import { IconAt, IconLock } from '@tabler/icons-react';

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Pingrate" },
        { name: "description", content: "Signup page for Pingrate" },
    ];
}

export default function Signup() {
    const [visible, { toggle }] = useDisclosure(false);
    return (
        <div>
            <p>Join Pingrate</p>
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
            />
            <Button variant="filled">Sign up</Button>
            <p>Do you have an account? Sign in</p>
        </div>
    );
}