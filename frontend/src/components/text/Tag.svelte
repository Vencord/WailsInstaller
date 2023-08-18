<!--
  Vencord WebInstaller, a web frontend for the Vencord Installer
  Copyright (c) 2023 Vendicated, Justice Almanzar and Vencord contributors
  SPDX-License-Identifier: GPL-3.0-or-later
-->

<script lang="ts">
    let className: string;
</script>

<span class="overline sm"><slot /></span>

<style>
    span {
        display: inline-block;
        background: var(--bg-0);
        border: 1px solid var(--bg-current-word);
        border-radius: 2px;
        padding: 0.25em 0.5rem;
        color: var(--color, var(--grey-1));
        font-family: var(--font);
        transform: translateY(-1px);
    }
</style>
