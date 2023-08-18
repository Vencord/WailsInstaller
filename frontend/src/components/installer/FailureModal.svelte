<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import Heading from "../text/Heading.svelte";
    export let message: string | null = null;
</script>

<section role="dialog">
    <Heading tag="h6" --color="var(--accent-red)">Oh no!</Heading>
    <p>Unable to do the thing u were trying to do :&lpar;&lpar;&lpar;&lpar;&lpar;</p>
    {#if message}
        <p>{message}</p>
    {/if}
</section>

<style>
    section {
        padding: 2em;
        text-align: center;
    }
</style>
