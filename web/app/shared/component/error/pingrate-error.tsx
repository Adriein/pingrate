import classes from "./error.module.css";

interface PingrateErrorProps {
    message: string
}

export default function PingrateError({ message }: PingrateErrorProps) {
    return (
        <div className={classes.errorCard}>
            <strong>Oh no!</strong>
            <p>{message}</p>
            <small>
                We've logged the issue and are working on a solution. Still having problems?
                Contact our support team.
            </small>
        </div>
    );
}