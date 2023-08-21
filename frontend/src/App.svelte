<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import { MinusIcon, XIcon } from "svelte-feather-icons";

    import { windowStore } from "./components/windows/index.js";
    import { launchers } from "./components/windows/launchers.js";

    import * as WailsRuntime from "../wailsjs/runtime/runtime.js";

    import VencordIcon from "./components/VencordIcon.svelte";
    import Window from "./components/windows/Window.svelte";
    import Launcher from "./components/windows/Launcher.svelte";

    $: windows = Object.values($windowStore);

    const autolaunch = +(localStorage.getItem("autolaunch") || "0") || 0;
    (launchers[autolaunch] ?? launchers[0]).onClick();
</script>

<div class="frame">
    {#await WailsRuntime.Environment() then env}
        {#if env.platform === "windows"}
            <div class="titlebar" role="application">
                <div class="icon">
                    <VencordIcon />
                </div>
                <div class="title body sm">Vencord Installer</div>
                <div class="spacer"></div>
                <div class="buttons">
                    <button title="Minimize" on:click={WailsRuntime.WindowMinimise}>
                        <MinusIcon size="1.5x" />
                    </button>
                    <button title="Close" class="close" on:click={WailsRuntime.Quit}>
                        <XIcon size="1.5x" />
                    </button>
                </div>
            </div>
        {/if}
    {/await}
    <div class="content">
        <div class="launchers">
            {#each launchers as launcher}
                <Launcher {launcher} />
            {/each}
        </div>
        {#each windows as window (window.props.id)}
            <Window {...window.props}>
                <svelte:component this={window.content} {...window.contentProps} />
            </Window>
        {/each}
    </div>
</div>

<style>
    .launchers {
        padding: 1em;
    }

    .frame {
        display: flex;
        flex-direction: column;
        position: absolute;
        overflow: hidden;
        outline: 1px solid var(--bg-3);

        top: 0;
        left: 0;
        width: 100vw;
        height: 100vh;
    }

    .titlebar {
        display: flex;
        align-items: center;
        background: #1e2021;
        color: #d4be98;
        cursor: default;
        height: 2rem;
        user-select: none;
    }

    .icon {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 1.25rem;
        height: 1.25rem;
        margin: 0 1rem;
    }

    .spacer {
        flex: 1;
    }

    .buttons {
        display: flex;
        height: 100%;
    }

    .buttons button {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 2.5rem;
        transition:
            background-color 0.2s ease-in-out,
            color 0.2s ease-in-out;
    }

    .buttons button:hover {
        background-color: #3c3836;
    }

    .buttons button:active {
        background-color: #141617;
    }

    .buttons button.close:hover {
        background-color: #ea6962;
        color: #1e2021;
    }

    .buttons button.close:active {
        background-color: #c14a4a;
        color: #1e2021;
    }

    .content {
        flex: 1;
    }

    .titlebar {
        --wails-draggable: drag;
    }

    .buttons {
        --wails-draggable: no-drag;
    }
</style>
