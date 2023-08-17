/*
 * Vencord WebInstaller, a web frontend for the Vencord Installer
 * Copyright (c) 2023 Vendicated, Justice Almanzar, and Vencord contributors
 * SPDX-License-Identifier: GPL-3.0-or-later
 */

import { svelte } from "@sveltejs/vite-plugin-svelte";
import { defineConfig } from "vite";

export default defineConfig({
    plugins: [svelte()],
    publicDir: "static",
    build: {
        outDir: "dist",
        emptyOutDir: true,
        sourcemap: true,
        minify: true
    }
});