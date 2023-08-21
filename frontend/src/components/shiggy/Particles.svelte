<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import { bgParticles, clickParticles } from "./particles";

    const clickParticlesStore = clickParticles.store;
    const clickParticlesTransition = clickParticles.transition;
    $: activeShiggies = $clickParticlesStore;

    const bgParticlesStore = bgParticles.store;
    const bgParticlesTransition = bgParticles.transition;
    $: passiveShiggies = $bgParticlesStore;

    const shiggyPng = "https://media.discordapp.net/stickers/1039992459209490513.png";
</script>

{#each passiveShiggies as shiggy (shiggy.id)}
    <div
        class="shiggy-bg"
        style="background-image:url({shiggyPng});left:calc({~~(shiggy.data.x * 100)}% - 1em);top:calc({shiggy.data
            .y}px - 1em);"
        out:bgParticlesTransition={shiggy.data}
    ></div>
{/each}

<slot />

{#each activeShiggies as shiggy (shiggy.id)}
    <div
        class="shiggy-clicked"
        style="background-image:url({shiggyPng});left:calc({shiggy.data.x}px - 1em);top:calc({shiggy.data.y}px - 1em);"
        out:clickParticlesTransition={shiggy.data}
    ></div>
{/each}

<style>
    .shiggy-clicked,
    .shiggy-bg {
        width: 2em;
        height: 2em;
        background-size: contain;
        pointer-events: none;
        position: absolute;
        opacity: 0;
    }
    .shiggy-bg {
        filter: blur(1px);
    }
</style>
