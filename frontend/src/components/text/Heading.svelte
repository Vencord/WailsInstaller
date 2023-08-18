<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import { getContext, setContext } from "svelte";

    export let tag: "h1" | "h2" | "h3" | "h4" | "h5" | "h6";

    const prevHeadingLevel = getContext<number>("headingLevel") ?? 0;
    const headingLevel = prevHeadingLevel + 1;
    setContext("headingLevel", headingLevel);
</script>

<svelte:element this={`h${headingLevel}`} class={`heading ${tag}`}><slot /></svelte:element>

<style>
    .heading {
        color: var(--color, inherit);
        margin: 0;
        font-weight: 600;
        letter-spacing: -2%;
    }
</style>
