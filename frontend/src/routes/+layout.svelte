<script lang="ts">
import { resolve } from '$app/paths';
import { onMount } from 'svelte';
import Header from './Header.svelte';
import { goto } from '$app/navigation';
import { setLocale, getLocale, locales } from '$lib/paraglide/runtime.js';
import {
	CheckDefinitionsExist,
	CheckForDefsUpdates,
	UpdateDefinitions,
	CheckForUpdates,
	ConfigExists,
	GetConfig,
	SaveConfig
} from '../../wailsjs/go/main/App';
import { updateAvailable } from '$lib/stores/updateStore';
import '../app.css';

type Theme = 'light' | 'dark' | 'triduilight' | 'triduidark' | 'auto';
type Settings = {
	theme?: Theme;
	language?: string;
	languageManuallySet?: boolean;
	autoUpdateDefinitions?: boolean;
	checkAppUpdatesOnStartup?: boolean;
	[key: string]: unknown;
};

let { children } = $props();

function getSetting<T = unknown>(key: string): T | null {
	const settingsData = localStorage.getItem('_trid_settings_');
	if (!settingsData) return null;
	const settings = JSON.parse(settingsData) as Record<string, T>;
	return settings[key] !== undefined ? (settings[key] as T) : null;
}

// function setSetting(key: string, value: any): void {
// 	const settings = JSON.parse(localStorage.getItem('_trid_settings_') || '{}');
// 	settings[key] = value;
// 	localStorage.setItem('_trid_settings_', JSON.stringify(settings));
// 	try {
// 		SaveSetting(key, value);
// 	} catch (err) {
// 		console.debug('SaveSetting call failed (backend may be unavailable):', err);
// 	}
// }

function applyTheme(theme: 'light' | 'dark' | 'triduilight' | 'triduidark' | 'auto') {
	if (typeof window === 'undefined') return;

	if (theme === 'auto') {
		const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
		document.documentElement.setAttribute('data-theme', prefersDark ? 'triduidark' : 'triduilight');
	} else {
		document.documentElement.setAttribute('data-theme', theme);
	}
}

onMount(() => {
	(async () => {
		try {
			const exists = await ConfigExists();
			if (exists) {
				try {
					const cfg = await GetConfig();
					if (cfg) {
						const parsed = JSON.parse(cfg) as Settings;
						localStorage.setItem('_trid_settings_', JSON.stringify(parsed));
						if (parsed.theme) applyTheme(parsed.theme);
						if (parsed.language && (locales as readonly string[]).includes(parsed.language)) {
							setLocale(parsed.language as typeof locales[number], { reload: false });
						}
					}
				} catch (err) {
					console.error('Failed to read config from backend:', err);
				}
			} else {
				try {
					const settingsData = localStorage.getItem('_trid_settings_') || '{}';
					await SaveConfig(settingsData);
					try {
						const parsed = JSON.parse(settingsData) as Settings;
						if (parsed.theme) applyTheme(parsed.theme);
						if (parsed.language && (locales as readonly string[]).includes(parsed.language)) {
							setLocale(parsed.language as typeof locales[number], { reload: false });
						}
					} catch (e) {
							console.debug('Failed to apply migrated settings', e);
					}
				} catch (err) {
					console.error('Failed to create config from localStorage:', err);
				}
			}
		} catch (err) {
			console.debug('Config sync skipped:', err);
		}
	})();

	const handleDragEnter = (e: DragEvent) => {
		if (e.dataTransfer?.types.includes('Files')) {
			e.preventDefault();
			if (window.location.pathname !== '/') goto(resolve('/'));
		}
	};

	const handleDragOver = (e: DragEvent) => {
		if (e.dataTransfer?.types.includes('Files')) e.preventDefault();
	};

	const handleDrop = (e: DragEvent) => {
		try {
			const tgt = e.target as Element | null;
			if (tgt && typeof tgt.closest === 'function' && tgt.closest('[data-app-drop-target]')) return;
		} catch {
			// ignore
		}
		if (e.dataTransfer?.types.includes('Files')) {
			e.preventDefault();
			e.stopPropagation();
		}
	};

	window.addEventListener('dragenter', handleDragEnter, true);
	window.addEventListener('dragover', handleDragOver, true);
	window.addEventListener('drop', handleDrop, true);

	const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
	const handleSystemThemeChange = () => {
		const currentTheme = getSetting('theme');
		if (currentTheme === 'auto' || !currentTheme) applyTheme('auto');
	};
	mediaQuery.addEventListener('change', handleSystemThemeChange);

		const hasManualPreference = getSetting('languageManuallySet');
	if (!hasManualPreference) {
		const browserLanguages = navigator.languages || [navigator.language];
		for (const browserLang of browserLanguages) {
			const langCode = browserLang.split('-')[0].toLowerCase();
			if ((locales as readonly string[]).includes(langCode)) {
				const currentLocale = getLocale();
				if (currentLocale !== langCode) setLocale(langCode as typeof locales[number], { reload: false });
				break;
			}
		}
	}

	const autoUpdateDefinitions = getSetting('autoUpdateDefinitions');
	if (autoUpdateDefinitions !== false) {
		CheckDefinitionsExist()
			.then((exists) => (exists ? CheckForDefsUpdates() : null))
			.then((updateInfo) => {
				if (updateInfo && !updateInfo.isUpToDate) return UpdateDefinitions();
			})
			.catch((err) => console.error('Auto-update definitions failed:', err));
	}

	const checkAppUpdatesOnStartup = getSetting('checkAppUpdatesOnStartup');
	if (checkAppUpdatesOnStartup !== false) {
		CheckForUpdates()
			.then((info) => {
				if (info) updateAvailable.set(info);
			})
			.catch((err) => console.error('Background app update check failed:', err));
	}

	return () => {
		window.removeEventListener('dragenter', handleDragEnter, true);
		window.removeEventListener('dragover', handleDragOver, true);
		window.removeEventListener('drop', handleDrop, true);
		mediaQuery.removeEventListener('change', handleSystemThemeChange);
	};
});
</script>

<div class="app h-screen flex flex-col">
	<Header />

	<main class="flex-1 overflow-auto">
		{@render children()}
	</main>
</div>

