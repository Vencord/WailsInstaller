<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import type { Launcher } from "./launchers";

    export let launcher: Launcher;
</script>

<button on:click={launcher.onClick}>
    <div class="icon"><svelte:component this={launcher.icon} /></div>
    <span class="body sm">{launcher.name}</span>
</button>

<style>
    button {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: 0.5em;
        width: 4em;
        text-align: center;
        border-radius: 0.5em;
        cursor: pointer;
    }
    button:hover {
        outline: 2px solid var(--bg-5);
        background-color: var(--bg-3);
    }
    .icon {
        width: 3em;
        height: 3em;
        margin: 0 auto;
    }
</style>
