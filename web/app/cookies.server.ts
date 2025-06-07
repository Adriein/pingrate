import {type Cookie, createCookie} from "react-router";

export const sessionCookie: Cookie = createCookie("$session", {
    httpOnly: true,
    maxAge: 604_800, // one week
});