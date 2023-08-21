/*
 * Vencord WebInstaller, a web frontend for the Vencord Installer
 * Copyright (c) 2023 Vendicated, Justice Almanzar, and Vencord contributors
 * SPDX-License-Identifier: GPL-3.0-or-later
 */

import { circOut } from "svelte/easing";
import { writable } from "svelte/store";
import type { TransitionConfig } from "svelte/transition";

type Transition<D = unknown> = (node: Element, data: D) => TransitionConfig;
type Particle<D = unknown> = {
    id: string;
    data: D;
};

function createParticleStore<D>(transition: Transition<D>) {
    const store = writable<Particle<D>[]>([]);

    return {
        store,
        transition,
        create(data: D) {
            store.update(particles => {
                const newId = Math.random().toString(36).slice(2, 9);
                particles.push({
                    id: newId,
                    data
                });
                delay(() => {
                    store.update(particles => particles.filter(({ id }) => id !== newId));
                }, 0);
                return particles;
            });
        }
    };
}

function delay(fn: (...args: any[]) => any, ms: number) {
    const start = Date.now();
    const frame = () => {
        if (Date.now() - start >= ms) {
            fn();
        } else {
            requestAnimationFrame(frame);
        }
    };
    requestAnimationFrame(frame);
}

export const clickParticles = createParticleStore<{
    x: number;
    y: number;
    v0: [number, number];
}>((node, { v0: [vx, vy] }) => {
    const g = 1;
    const s = 20;
    const vs = 2;
    return {
        duration: 1000,
        easing: circOut,
        css: t => {
            const q = 1 - 2 * Math.abs(t - 0.5);
            const rt = 1 - t;
            const x = rt * vx * vs;
            const y = rt * vy * vs - g * rt * rt;
            return `opacity:${q};transform:translate(${x * s}px,${-y * s}px)`;
        }
    };
});

export const bgParticles = createParticleStore<{
    x: number;
    y: number;
    v: number;
    s: number;
}>((node, { v, s }) => {
    const dy = v * 400 + 40;
    return {
        duration: dy * 10,
        // easing: circOut,
        css: t => {
            const q = 1 - 2 * Math.abs(t - 0.5);
            const rt = 1 - t;
            return `opacity:${q * 0.5};transform:scale(${s}) translateY(${dy * rt}px)`;
        }
    };
});
