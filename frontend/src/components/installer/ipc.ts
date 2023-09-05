import FailureModal from "./FailureModal.svelte";
import SuccessModal from "./SuccessModal.svelte";

import { openWindow } from "../windows";

import * as Installer from "../../../wailsjs/go/installer/Installer";

export type IPCCall = (typeof Installer)["Patch" | "Repair" | "Unpatch" | "InstallOpenAsar" | "UninstallOpenAsar"];

export function getOpPastTense(op: IPCCall) {
    if (op === Installer.Patch) return "installed";
    if (op === Installer.Repair) return "repaired";
    if (op === Installer.Unpatch) return "uninstalled";
    if (op === Installer.InstallOpenAsar) return "installed OpenAsar";
    if (op === Installer.UninstallOpenAsar) return "uninstalled OpenAsar";
    return "did something?";
}

export async function doAction(op: IPCCall, path: string, onAction: () => void) {
    await op(path)
        .then(() => {
            openWindow(
                SuccessModal,
                {
                    verb: getOpPastTense(op)
                },
                {
                    title: "Woohoo!",
                    width: 400,
                    height: 510
                }
            );
        })
        .catch(error => {
            const message = typeof error === "string" ? error : null;
            if (!message) console.error(error);
            openWindow(
                FailureModal,
                {
                    message,
                    path,
                    op,
                    onAction
                },
                {
                    title: "Uh oh!",
                    width: 450,
                    height: 200
                }
            );
        })
        .finally(() => {
            onAction?.();
        });
}
