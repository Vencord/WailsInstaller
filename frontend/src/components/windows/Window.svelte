<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import type { SvelteComponent } from "svelte";
    import { cubicInOut } from "svelte/easing";
    import type { TransitionConfig } from "svelte/transition";
    import { CircleIcon, XIcon } from "svelte-feather-icons";
    import type { Constructor } from "type-fest";

    import { closeWindow, getFocusZIndex } from ".";

    export let title: string;
    export let id: string;
    export let icon: Constructor<SvelteComponent> | null = null;

    export let backgroundColor = "var(--bg-1)";

    export let width: number;
    export let height: number;
    export let minWidth: number | null = width;
    export let minHeight: number | null = height;

    export let maximized = false;
    export let closable = true;

    function close() {
        closeWindow(id);
    }

    function clamp(x: number, y: number): [x: number, y: number] {
        return [
            Math.max(Math.min(x, window.innerWidth - width), 0),
            Math.max(Math.min(y, window.innerHeight - height), 32)
        ];
    }

    // Dragging
    let dragAnchor: [number, number] | null = null;
    let x = Math.floor((window.innerWidth - width) / 2);
    let y = Math.floor((window.innerHeight - height) / 2);
    function onDragStart(event: MouseEvent) {
        dragAnchor = [event.x - x, event.y - y];
    }
    function onDrag(event: MouseEvent) {
        if (!dragAnchor || (!event.x && !event.y)) return;
        const [anchorX, anchorY] = dragAnchor;
        [x, y] = clamp(event.x - anchorX, event.y - anchorY);
        window.getSelection?.()?.empty?.();
    }
    function onDragEnd(event: MouseEvent) {
        if (dragAnchor) {
            const [anchorX, anchorY] = dragAnchor;
            [x, y] = clamp(event.x - anchorX, event.y - anchorY);
        }
        dragAnchor = null;
    }

    // Focusing
    let z = 0;
    function onFocus() {
        z = getFocusZIndex();
    }
    onFocus();

    // Resizing
    function resize(node: HTMLElement) {
        const observer = new ResizeObserver(entries => {
            ({ width, height } = entries[0].contentRect);
        });
        observer.observe(node);
        return {
            destroy() {
                observer.disconnect();
            }
        };
    }

    // Inline styles
    let style: string;
    $: {
        const styleProps: Record<string, string> = {
            "background-color": backgroundColor
        };
        if (!maximized) {
            styleProps.width = `${width}px`;
            styleProps.height = `${height}px`;
            styleProps.left = `${x}px`;
            styleProps.top = `${y}px`;
            styleProps["z-index"] = `${z}`;
        }
        if (minHeight) styleProps["min-height"] = `${minHeight}px`;
        if (minWidth) styleProps["min-width"] = `${minWidth}px`;
        style = Object.entries(styleProps)
            .map(([key, value]) => `${key}:${value}`)
            .join(";");
    }

    function transition(node: Element): TransitionConfig {
        return {
            duration: 200,
            easing: cubicInOut,
            css: t => `transform: scale(${t / 2 + 0.5}); opacity: ${t}`
        };
    }
</script>

<svelte:window on:mousemove={onDrag} on:mouseup={onDragEnd} />

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="frame" class:maximized use:resize on:mousedown={onFocus} {style} in:transition out:transition>
    <!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
    <div class="titlebar" role="application" on:mousedown={onDragStart}>
        <div class="icon">
            {#if icon}
                <svelte:component this={icon} />
            {/if}
        </div>
        <div class="title body sm">{title}</div>
        <div class="spacer"></div>
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <div class="buttons" on:mousedown={event => event.stopPropagation()}>
            <button on:click={close} title="Close"><XIcon size="1.5x" /></button>
        </div>
    </div>
    <div class="content">
        <slot />
    </div>
</div>

<style>
    .frame {
        display: flex;
        flex-direction: column;
        position: absolute;
        resize: both;
        overflow: hidden;
    }

    .frame:not(.maximized) {
        border-radius: 0.5rem;
        outline: 1px solid #3c3836;
        box-shadow:
            0px 2px 10.6px 0px rgba(0, 0, 0, 0.37),
            0px 16px 32px 0px rgba(0, 0, 0, 0.37);
    }
    .frame::-webkit-resizer {
        display: none;
    }

    .frame.maximized {
        top: var(--titlebar-height);
        left: 0;
        width: 100vw;
        height: calc(100vh - var(--titlebar-height));
        resize: none;
    }

    .titlebar {
        display: flex;
        align-items: center;
        color: #d4be98;
        cursor: default;
        height: 2rem;
        user-select: none;
    }

    .frame:not(.maximized) .titlebar {
        background: #1e2021;
    }

    .icon {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 1.25rem;
        height: 1.25rem;
        margin: 0 1rem;
    }
    .frame.maximized .icon,
    .frame.maximized .title {
        opacity: 0;
    }

    .spacer {
        flex: 1;
    }

    .buttons {
        display: flex;
        transition: opacity 0.2s ease-in-out;
        height: 100%;
    }

    .frame.maximized .buttons {
        opacity: 0;
    }
    .titlebar:hover .buttons {
        opacity: 1;
    }

    .buttons button {
        display: flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;
        width: 2.5rem;
    }

    .buttons button:hover {
        background-color: #ea6962;
        color: #1e2021;
    }

    .content {
        flex: 1;
    }
</style>
