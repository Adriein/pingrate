import classes from "./error.module.css";
import {Anchor, type MantineTheme, useMantineTheme} from "@mantine/core";
import React from "react";

interface PingrateErrorProps {
    message: string
}

export default function PingrateError({ message }: PingrateErrorProps) {
    const theme: MantineTheme = useMantineTheme();
    return (
        <div className={classes.errorCard}>
            <strong>Oh no!</strong>
            <p>{message}</p>
            <small>
                We've logged the issue and are working on a solution. Still having problems? {' '}
                <Anchor
                    size="xs"
                    styles={{
                        root: {
                            color: theme.colors.pingrateAccent[10]
                        }
                    }}
                >
                    Contact our support team.
                </Anchor>
            </small>
        </div>
    );
}