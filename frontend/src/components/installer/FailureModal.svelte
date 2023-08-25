<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import { Environment, BrowserOpenURL } from "../../../wailsjs/runtime/runtime.js";
    import * as Installer from "../../../wailsjs/go/installer/Installer.js";
    import SuccessModal from "./SuccessModal.svelte";
    import Heading from "../text/Heading.svelte";
    import Button from "../Button.svelte";
    import { closeWindow, openWindow } from "../windows/index.js";

    type IPCCall = (typeof Installer)["Patch" | "Repair" | "Unpatch" | "InstallOpenAsar" | "UninstallOpenAsar"];

    export let message: string | null = null;
    export let path: string = "";
    export let op: IPCCall;
    export let getOpPastTense: (IPCCall) => string;
    export let onAction: () => void;

    async function runAndShowSuccess() {
        await op(path);
        openWindow(
            SuccessModal,
            {
                verb: getOpPastTense(op)
            },
            {
                title: "Success",
                width: 400,
                height: 510
            }
        );
        // closeWindow(id);
        onAction();
    }

    let buttonPress: () => void;

    function waitForButton() {
        return new Promise<void>(res => {
            buttonPress = res;
        });
    }
</script>

<section role="dialog">
    <Heading tag="h6" --color="var(--accent-red)">Oh no!</Heading>
    <p>Unable to do the thing u were trying to do :&lpar;&lpar;&lpar;&lpar;&lpar;</p>
    {#if message}
        {#if message.includes("file exists") || message.includes("permission denied")}
            {#await Installer.CheckForOwnership(path)}
                <p>{message}</p>
            {:then isOwned}
                {#if !isOwned}
                    {#await Environment()}
                        <p>{message}</p>
                    {:then env}
                        {#if env.platform === "darwin"}
                            {#await waitForButton()}
                                <p>Please allow vencord to fix discord's permissions</p>
                                <Button on:click={buttonPress}>Try again as admin</Button>
                            {:then}
                                {#await Installer.PromptForChown(path)}
                                    <p>Please allow vencord to fix discord's permissions</p>
                                    <Button on:click={() => {}}>Try again as admin</Button>
                                {:then}
                                    {#await runAndShowSuccess()}
                                        <p>Working...</p>
                                    {:then}
                                        <p>Success! You can close this window</p>
                                    {:catch}
                                        <p>Please grant the installer full disk access</p>
                                        <Button
                                            on:click={() =>
                                                BrowserOpenURL(
                                                    "x-apple.systempreferences:com.apple.preference.security?Privacy_AllFiles"
                                                )}>Open Security Settings</Button
                                        >
                                    {/await}
                                {:catch error}
                                    <p>{error}</p>
                                {/await}
                            {/await}
                        {:else}
                            <p>{message}</p>
                        {/if}
                    {/await}
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
