<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    import PatchSelect from "../input/PatchSelect.svelte";
    import Tag from "../text/Tag.svelte";
    import Actions from "./Actions.svelte";

    import * as Installer from "../../../wailsjs/go/installer/Installer";
    import type { installer } from "../../../wailsjs/go/models";
    import ShiggyIcon from "../ShiggyIcon.svelte";

    const userDataDir = "%APPDATA%\\Vencord";

    const branches = ["stable", "canary", "ptb", "development"];
    let installPromise: Promise<installer.DiscordData[]>;

    function fetchInstalls() {
        const promise = Installer.ListInstalls().then(installs => {
            return installs.sort((a, b) => branches.indexOf(a.branch) - branches.indexOf(b.branch));
        });

        if (!installPromise) installPromise = promise;
        else promise.then(installs => (installPromise = Promise.resolve(installs)));
    }
    fetchInstalls();

    let selectedInstall = "";

    let isOpenAsarSelected = false;
    $: {
        installPromise
            .then(installs => {
                const install = installs.find(e => e.path === selectedInstall);
                isOpenAsarSelected = install?.isOpenAsar ?? false;
            })
            .catch(() => {
                isOpenAsarSelected = false;
            });
    }

    const currentVersion = "v1.2.6";
    const latestVersion = "v1.2.7";

    let busy = false;
</script>

<section>
    <div>
        <p>
            Files will be downloaded to <Tag>{userDataDir}</Tag>.
        </p>
        <p>
            To customize this location, set the environment variable <Tag>VENCORD_USER_DATA_DIR</Tag> and restart the installer.
        </p>
        <hr />
        <div class="installer">
            <div class="version">
                <p class="overline sm current">Current</p>
                <Tag>{currentVersion}</Tag>
            </div>
            <div class="version">
                <p class="overline sm latest">latest</p>
                <Tag>{latestVersion}</Tag>
            </div>
        </div>
        <hr />
        {#await installPromise}
            <p>Loading installs...</p>
        {:then installs}
            <PatchSelect options={installs} bind:selected={selectedInstall} />
        {:catch error}
            <p>Failed to load installs: {String(error)}</p>
        {/await}
    </div>

    <Actions bind:busy path={selectedInstall} onAction={fetchInstalls} isOpenAsar={isOpenAsarSelected} />

    {#if busy}
        <div class="loader">
            <div>
                <ShiggyIcon/>
                <h1>Busy...</h1>
                <b>(shiggy's doing her best)</b>
            </div>
        </div>
    {/if}
</section>

<style>
    section {
        display: flex;
        flex-direction: column;
        gap: 2rem;
    }

    .installer {
        display: flex;
        gap: 1rem;
        align-items: center;
        justify-content: center;
    }
    .version {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 0.1rem;
    }
    .current {
        color: var(--accent-purple);
    }
    .latest {
        color: var(--accent-green);
    }

    .loader {
        display: flex;

        position: absolute;

        inset: 0;

        background: rgba(0, 0, 0, 0.4);

        justify-content: center;
        align-items: center;
        text-align: center;

        padding: 16rem;
        box-sizing: border-box;
    }

    .loader div {
        background: var(--bg-0);

        margin-top: 1rem;

        padding: 2rem;
        border-radius: 0.5rem;
        box-shadow:
            0px 2px 10.6px 0px rgba(0, 0, 0, 0.37),
            0px 16px 32px 0px rgba(0, 0, 0, 0.37);
    }

    .loader h1 {
        margin-top: 0rem;
        margin-bottom: 1rem;

        font-size: 3rem;
    }

    .loader b {
        color: var(--grey-1);
    }
</style>
