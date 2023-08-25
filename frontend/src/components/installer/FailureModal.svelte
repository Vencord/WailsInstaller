<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import { Environment, BrowserOpenURL } from "../../../wailsjs/runtime/runtime.js";
    import Heading from "../text/Heading.svelte";
    import Button from "../Button.svelte";
    export let message: string | null = null;
</script>

<section role="dialog">
    <Heading tag="h6" --color="var(--accent-red)">Oh no!</Heading>
    <p>Unable to do the thing u were trying to do :&lpar;&lpar;&lpar;&lpar;&lpar;</p>
    {#if message}
        {#if message.includes("operation not permitted")}
            {#await Environment()}
                <p>{message}</p>
            {:then env}
                {#if env.platform === "darwin"}
                    <p>Please grant the installer full disk access</p>
                    <Button
                        on:click={() =>
                            BrowserOpenURL("x-apple.systempreferences:com.apple.preference.security?Privacy_AllFiles")}
                        >Open Security Settings</Button
                    >
                {:else}
                    <p>{message}</p>
                {/if}
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
