<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import { openWindow } from "../windows";
    import { clickParticles } from "./particles";
    import Particles from "./Particles.svelte";
    import { addPoints, pointStore, shiggingStore } from "./shiggy";
    import Shop from "./Shop.svelte";

    $: points = $pointStore;
    $: shigging = $shiggingStore;

    function shig(event: MouseEvent) {
        console.log(event);
        const angle = Math.random() * Math.PI;
        clickParticles.create({
            x: event.offsetX,
            y: event.offsetY,
            v0: [Math.cos(angle), Math.sin(angle)]
        });
        addPoints(1);
    }

    function openShop(event: MouseEvent) {
        event.stopPropagation();
        openWindow(
            Shop,
            {},
            {
                title: "Shiggy Shop",
                width: 400,
                height: 600
            }
        );
    }
</script>

<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
<!-- svelte-ignore a11y-click-events-have-key-events -->
<main on:click={shig}>
    <Particles>
        <p class="points overline lg">{points} shiggies</p>
        <div class="shig" class:shigging></div>
    </Particles>
    <div class="spacer" />
    {#if points >= 100}
        <button class="shop" on:click={openShop}>Open shop?</button>
    {/if}
</main>

<style>
    main {
        padding: 1em;
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 2em;
        height: 100%;
        box-sizing: border-box;
        user-select: none;
        position: relative;
        overflow: hidden;
        cursor: pointer;
    }
    .spacer {
        flex: 1;
    }
    .shig {
        width: 8em;
        height: 8em;
        background: url(https://media.discordapp.net/stickers/1039992459209490513.png?passthrough=false),
            url(https://media.discordapp.net/stickers/1039992459209490513.png?passthrough=true);
        background-size:
            contain,
            0 0;
        position: relative;
    }
    .shig.shigging {
        background-size:
            0 0,
            contain;
    }

    .shop {
        color: var(--accent-green);
        text-decoration: underline;
        cursor: pointer;
    }

    .shig,
    .points {
        pointer-events: none;
    }
</style>
