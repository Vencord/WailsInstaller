<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import Heading from "../text/Heading.svelte";

    export let verb: string;
    // export let path: string;
</script>

<section role="dialog">
    <Heading tag="h6" --color="var(--accent-green)">Successfully {verb}!</Heading>
    <p>Enjoy this shiggy:</p>
    <img src="https://media.discordapp.net/stickers/1039992459209490513.png?size=240" alt="shiggy" />
</section>

<style>
    section {
        padding: 2em;
        text-align: center;
    }
</style>
