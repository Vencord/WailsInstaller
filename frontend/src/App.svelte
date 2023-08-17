<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import { MinusIcon, XIcon } from "svelte-feather-icons";

    import * as WailsRuntime from "../wailsjs/runtime/runtime.js";

    import VencordIcon from "./components/VencordIcon.svelte";
</script>

<div class="frame">
    {#await WailsRuntime.Environment() then env}
        {#if env.platform === "windows"}
            <div class="titlebar" role="application">
                <div class="icon">
                    <VencordIcon />
                </div>
                <div class="title body sm">Vencord Installer Shell</div>
                <div class="spacer"></div>
                <div class="buttons">
                    <button title="Minimize" on:click={WailsRuntime.WindowMinimise}>
                        <MinusIcon size="1.1x" />
                    </button>
                    <button title="Close" class="close" on:click={WailsRuntime.Quit}>
                        <XIcon size="1.1x" />
                    </button>
                </div>
            </div>
        {/if}
    {/await}
    <div class="content">
        <iframe src="https://beta.install.vencord.dev"></iframe>
    </div>
</div>

<style>
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
    }

    .buttons button:hover {
        background-color: #3c3836;
    }
    
    .buttons button.close:hover {
        background-color: #ea6962;
        color: #1e2021;
    }

    .content {
        flex: 1;
    }

    .content iframe {
        width: 100%;
        height: 100%;
        border: none;
    }

    .titlebar {
        --wails-draggable: drag;
    }

    .buttons {
        --wails-draggable: no-drag;
    }
</style>
