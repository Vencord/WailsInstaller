<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import type { HTMLButtonAttributes } from "svelte/elements";

    interface $$Props extends HTMLButtonAttributes {
        "--accent"?: string;
    }
</script>

<button {...$$props} on:click class={`label md ${$$props.class || ""}`}>
    <slot />
</button>

<style>
    button {
        text-align: center;
        box-sizing: border-box;
        padding: 0.75rem 1rem;
        border-radius: 0.75rem;
        background-color: var(--accent, var(--fg-1));
        color: var(--bg-current-word);
        cursor: pointer;
    }
    button:disabled {
        opacity: 0.7;
        cursor: default;
    }
</style>
