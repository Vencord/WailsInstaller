/*
 * Vencord WebInstaller, a web frontend for the Vencord Installer
 * Copyright (c) 2023 Vendicated, Justice Almanzar, and Vencord contributors
 * SPDX-License-Identifier: GPL-3.0-or-later
 */

import type { SvelteComponent } from "svelte";
import type { Constructor } from "type-fest";

import ShiggyIcon from "../ShiggyIcon.svelte";
import { openWindow } from ".";
import ShiggyGame from "../shiggy/ShiggyGame.svelte";

export type Launcher = {
    name: string;
    icon: Constructor<SvelteComponent>;
    onClick: () => void;
};

export const launchers: Launcher[] = [
    {
        name: "Shiggy Clicker",
        icon: ShiggyIcon,
        onClick: () => {
            openWindow(
                ShiggyGame,
                {},
                {
                    title: "Shiggy Clicker",
                    icon: ShiggyIcon,
                    width: 400,
                    height: 500
                }
            );
        }
    }
];
