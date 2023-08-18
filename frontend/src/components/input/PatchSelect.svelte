<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import { CheckIcon } from "svelte-feather-icons";

    import Heading from "../text/Heading.svelte";
    import Tag from "../text/Tag.svelte";

    import type { installer } from "../../../wailsjs/go/models";

    export let options: installer.DiscordData[];
    export let selected = options[0].path;

    const id = Math.floor(Math.random() * 100);
    const getBranchColor = (branch: string) => {
        const varName =
            {
                stable: "--accent-blue",
                ptb: "--accent-aqua",
                canary: "--accent-yellow",
                dev: "--grey-1"
            }[branch] ?? "--grey-1";
        return `var(${varName})`;
    };
</script>

<div role="radiogroup" class="container" aria-labelledby="label-patch" id={`group-${id}`}>
    <div class="legend" id="label-patch">
        <Heading tag="h6" --color="var(--accent-orange)">Please select an install to patch</Heading>
    </div>
    {#each options as { branch, path, isPatched, isOpenAsar }}
        <div class="row" class:selected={selected === path}>
            <div class="radio-wrap">
                <input type="radio" id={path} bind:group={selected} value={path} />
                <span class="radio"><div class="check" /></span>
            </div>
            <label for={path}>
                {#if isPatched}
                    <Tag --color="var(--accent-green)">
                        <span>PATCHED</span>
                        <CheckIcon size="1x" strokeWidth={5} class="dy-1" />
                    </Tag>
                {/if}
                {#if isOpenAsar}
                    <Tag>OpenAsar</Tag>
                {/if}
                {path}
                {#if branch}
                    <Tag --color={getBranchColor(branch)}>{branch}</Tag>
                {/if}
            </label>
        </div>
    {/each}
</div>

<style>
    .container {
        display: flex;
        gap: 0.75rem;
        flex-direction: column;
    }

    input[type="radio"],
    label {
        cursor: pointer;
    }

    .row {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        color: var(--fg-0);
    }
    .selected {
        color: var(--fg-1);
    }

    .radio-wrap {
        position: relative;
    }
    .radio {
        position: absolute;
        inset: 0;
        background-color: var(--bg-0);
        border-radius: 50%;
        pointer-events: none;
        box-sizing: border-box;
        border: 1px solid var(--bg-3);
        padding: 3px;
        display: flex;
    }
    .radio .check {
        background: var(--accent-purple);
        opacity: 0;
        height: 100%;
        width: 100%;
        border-radius: 50%;
    }

    input[type="radio"]:checked ~ .radio .check {
        opacity: 1;
    }
    .legend {
        margin-bottom: 0.5rem;
    }
    input[type="radio"] {
        display: block;
        margin: 0;
        width: 1.25rem;
        height: 1.25rem;
        opacity: 0;
    }
</style>
