/*
 * Vencord WebInstaller, a web frontend for the Vencord Installer
 * Copyright (c) 2023 Vendicated, Justice Almanzar, and Vencord contributors
 * SPDX-License-Identifier: GPL-3.0-or-later
 */

import "./app.css";

import App from "./App.svelte";

const app = new App({
    target: document.getElementById("app")!
});
export default app;
