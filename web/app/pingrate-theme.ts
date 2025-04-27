import {
    colorsTuple,
    DEFAULT_THEME,
    type DefaultMantineColor,
    type MantineColorsTuple, type MantineTheme,
} from "@mantine/core";

type PingrateExtendedColors =
    | 'pingratePrimary'
    | 'pingrateSecondary'
    | 'pingrateAccent'
    | 'pingrateBackground'

type ExtendedCustomColors =
    | PingrateExtendedColors
    | DefaultMantineColor;

declare module '@mantine/core' {
    export interface MantineThemeColorsOverride {
        colors: Record<ExtendedCustomColors, MantineColorsTuple>;
    }
}

export const PingrateTheme: MantineTheme = {
    ...DEFAULT_THEME,
    colors: {
        ...DEFAULT_THEME.colors,
        /** Index 10 is the original  */
        pingratePrimary: [
            "#fffbe2",
            "#fff6cd",
            "#ffeb9c",
            "#ffe066",
            "#ffd73a",
            "#ffd120",
            "#ffce10",
            "#e3b600",
            "#caa100",
            "#ae8b00",
            "#E6B800"
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
    } as Record<ExtendedCustomColors, MantineColorsTuple>,
}