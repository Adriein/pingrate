import {
    colorsTuple,
    DEFAULT_THEME,
    type MantineColorsTuple,
    type MantineTheme
} from "@mantine/core";

export const PingrateTheme: MantineTheme = {
    ...DEFAULT_THEME,
    colors: {
        ...DEFAULT_THEME.colors,
        /** Index 10 is the original  */
        pingratePrimary: [
            "#fffae0",
            "#fff3ca",
            "#ffe699",
            "#ffd863",
            "#ffcd36",
            "#ffc518",
            "#ffc102",
            "#e3aa00",
            "#ca9700",
            "#af8200",
            "#FFCC33",
        ] as MantineColorsTuple,
        /** Index 10 is the original  */
        pingrateSecondary: [
            "#f5f5f5",
            "#e7e7e7",
            "#cdcdcd",
            "#b2b2b2",
            "#9a9a9a",
            "#8b8b8b",
            "#848484",
            "#717171",
            "#656565",
            "#575757",
            "#3C3C3C",
        ] as MantineColorsTuple,
        /** Index 10 is the original  */
        pingrateAccent: [
            "#fff6e1",
            "#ffeccb",
            "#ffd79a",
            "#ffc164",
            "#ffaf37",
            "#ffa31b",
            "#ff9d09",
            "#e38800",
            "#cb7800",
            "#b06700",
            "#FF9900"
        ] as MantineColorsTuple,
        pingrateBackground: colorsTuple('#F9F9F9'),
    },
}