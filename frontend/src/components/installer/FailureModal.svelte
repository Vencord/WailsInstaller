<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import * as Installer from "../../../wailsjs/go/installer/Installer.js";
    import SuccessModal from "./SuccessModal.svelte";
    import Heading from "../text/Heading.svelte";
    import Button from "../Button.svelte";
    import { closeWindow, openWindow } from "../windows/index.js";
    import { BrowserOpenURL } from "../../../wailsjs/runtime/runtime.js";

    type IPCCall = (typeof Installer)["Patch" | "Repair" | "Unpatch" | "InstallOpenAsar" | "UninstallOpenAsar"];

    export let message: string | null = null;
    export let path: string = "";
    export let op: IPCCall;
    export let getOpPastTense: (IPCCall) => string;
    export let onAction: () => void;
    export let _windowId: string;

    async function runAndShowSuccess() {
        if (!await Installer.PromptForChown(path)) return;

        closeWindow(_windowId);

        setTimeout(async () => {
            await op(path);
            openWindow(
                SuccessModal,
                {
                    verb: getOpPastTense(op)
                },
                {
                    title: "Woohoo!",
                    width: 400,
                    height: 510
                }
            );
            onAction();
        });
    }
</script>

<section role="dialog">
    <Heading tag="h6" --color="var(--accent-red)">Oh no!</Heading>
    <p>Something went wrong!</p>
    {#if message}
        {#if message.includes("file exists") || message.includes("permission denied")}
            {#await Installer.CheckForOwnershipDarwin(path)}
                <p>{message}</p>
            {:then isOwned}
                {#if !isOwned}
                    <p>Hmm... seems like you've encountered a Mac-specific problem! Usually, this is one of two things:</p>
                    <p>
                        Your Discord installation's permissions appear to be broken. Luckily, we offer a simple tool to
                        fix this.
                    </p>
                    <Button on:click={runAndShowSuccess}>Repair Permissions</Button>
                    <p>
                        Sometimes the installer needs Full Disk Access, though usually the above should suffice.
                    </p>
                    <Button on:click={() =>
                        BrowserOpenURL("x-apple.systempreferences:com.apple.preference.security?Privacy_AllFiles")}>
                        Open Security Settings
                    </Button>
                {:else}
                    <p>{message}</p>
                {/if}
            {:catch error}
                <p>Unable to check for root: {error}</p>
            {/await}
        {:else}
            <p>{message}</p>
        {/if}
    {/if}
</section>

<style>
    section {
        padding: 2em;
        text-align: center;
    }
</style>
