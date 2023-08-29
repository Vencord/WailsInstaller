<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import { openWindow } from "../windows";
    import Button from "../Button.svelte";
    import FailureModal from "./FailureModal.svelte";
    import SuccessModal from "./SuccessModal.svelte";

    import * as IPC from "./ipc";

    import * as Installer from "../../../wailsjs/go/installer/Installer";

    export let path: string;
    export let isOpenAsar: boolean = false;

    export let onAction: () => void;

    export let busy = false;

    async function doAction(op: IPC.IPCCall) {
        busy = true;
        await IPC.doAction(op, path, onAction);
        busy = false;
    }
</script>

<section>
    <Button disabled={busy} --accent="var(--accent-green)" on:click={() => doAction(Installer.Patch)}>Install</Button>
    <Button disabled={busy} --accent="var(--accent-blue)" on:click={() => doAction(Installer.Repair)}>Repair</Button>
    <Button disabled={busy} --accent="var(--accent-red)" on:click={() => doAction(Installer.Unpatch)}>Uninstall</Button>
    {#if isOpenAsar}
        <Button disabled={busy} class="openasar" on:click={() => doAction(Installer.UninstallOpenAsar)}>
            Uninstall OpenAsar
        </Button>
    {:else}
        <Button disabled={busy} class="openasar" on:click={() => doAction(Installer.InstallOpenAsar)}>Install OpenAsar</Button>
    {/if}
</section>

<style>
    section {
        display: grid;
        grid-template-columns: repeat(4, minmax(0, 1fr));
        gap: 1rem;
        width: 100%;
    }

    section :global(.openasar) {
        background-color: var(--bg-3);
        border: 1px solid var(--fg-0);
        color: var(--fg-1);
    }
</style>
