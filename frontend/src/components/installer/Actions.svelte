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

    import * as Installer from "../../../wailsjs/go/installer/Installer";

    type IPCCall = (typeof Installer)["Patch" | "Repair" | "Unpatch" | "InstallOpenAsar" | "UninstallOpenAsar"];

    export let path: string;
    export let isOpenAsar: boolean = false;

    export let onAction: () => void;

    export let busy = false;

    function getOpPastTense(op: IPCCall) {
        if (op === Installer.Patch) return "installed";
        if (op === Installer.Repair) return "repaired";
        if (op === Installer.Unpatch) return "uninstalled";
        if (op === Installer.InstallOpenAsar) return "installed OpenAsar";
        if (op === Installer.UninstallOpenAsar) return "uninstalled OpenAsar";
        return "did something?";
    }

    async function doAction(op: IPCCall) {
        busy = true;
        await op(path)
            .then(() => {
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
            })
            .catch(error => {
                const message = typeof error === "string" ? error : null;
                if (!message) console.error(error);
                openWindow(
                    FailureModal,
                    {
                        message,
                        path,
                        op,
                        getOpPastTense,
                        onAction
                    },
                    {
                        title: "Uh oh!",
                        width: 450,
                        height: 200
                    }
                );
            })
            .finally(() => {
                busy = false;
                onAction?.();
            });
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
