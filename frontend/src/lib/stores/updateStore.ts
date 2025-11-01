import { writable } from 'svelte/store';
import type { main } from '../../../wailsjs/go/models';

export const updateAvailable = writable<main.UpdateInfo | null>(null);
