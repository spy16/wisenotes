export * from "./types";

import { persisted } from 'svelte-persisted-store'

import type { Profile } from "./types"

export async function getProfiles() {
    return [
        {
            id: "work",
            name: "Work",
            emoji: "👔",
        },
    ] as Profile[];
}

export const appConf = persisted('preferences', {
    theme: 'cupcake',
})