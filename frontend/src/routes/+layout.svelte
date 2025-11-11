<script lang="ts">
	import { onMount } from 'svelte';
	import Header from './Header.svelte';
	import { page } from '$app/stores';
	import { setLocale, getLocale, locales, baseLocale } from '$lib/paraglide/runtime.js';
	import { CheckDefinitionsExist, CheckForDefsUpdates, UpdateDefinitions, CheckForUpdates } from '../../wailsjs/go/main/App';
	import { updateAvailable } from '$lib/stores/updateStore';
	import '../app.css';
	let { children } = $props();

	function getSetting(key: string): any {
		const settingsData = localStorage.getItem('_trid_settings_');
		if (!settingsData) return null;
		
		const settings = JSON.parse(settingsData);
		return settings[key] !== undefined ? settings[key] : null;
	}

	function setSetting(key: string, value: any): void {
		const settings = JSON.parse(localStorage.getItem('_trid_settings_') || '{}');
		settings[key] = value;
		localStorage.setItem('_trid_settings_', JSON.stringify(settings));
	}

	// Apply theme immediately before mount to prevent flash
	if (typeof window !== 'undefined') {
		const savedTheme = getSetting('theme') as 'light' | 'dark' | 'auto' | 'triduidark' | 'triduilight' | null;
		const theme = savedTheme || 'auto';
		
		if (theme === 'auto') {
			const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
			document.documentElement.setAttribute('data-theme', prefersDark ? 'triduidark' : 'triduilight');
		} else {
			document.documentElement.setAttribute('data-theme', theme);
		}
	}

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
		// Listen for system theme changes when in auto mode
		const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
		const handleSystemThemeChange = () => {
			const currentTheme = getSetting('theme');
			if (currentTheme === 'auto' || !currentTheme) {
				applyTheme('auto');
			}
		};
		mediaQuery.addEventListener('change', handleSystemThemeChange);

		// Check if user has manually set a language preference
		const hasManualPreference = getSetting('languageManuallySet');
		
		if (!hasManualPreference) {
			// Auto-detect browser language on first visit
			const browserLanguages = navigator.languages || [navigator.language];
			
			for (const browserLang of browserLanguages) {
				// Extract base language code (e.g., 'en' from 'en-US')
				const langCode = browserLang.split('-')[0].toLowerCase();
				
				// Check if this language is available in our app
				if (locales.includes(langCode as any)) {
					const currentLocale = getLocale();
					// Only set if different from current
					if (currentLocale !== langCode) {
						setLocale(langCode as any, { reload: false });
					}
					break;
				}
			}
		}

		// Auto-update definitions on startup if enabled
		const autoUpdateDefinitions = getSetting('autoUpdateDefinitions');
		if (autoUpdateDefinitions !== false) { // Default to true if not set
			CheckDefinitionsExist()
				.then((exists) => {
					if (exists) {
						return CheckForDefsUpdates();
					}
				})
				.then((updateInfo) => {
					if (updateInfo && !updateInfo.isUpToDate) {
						return UpdateDefinitions();
					}
				})
				.catch((err) => {
					console.error('Auto-update definitions failed:', err);
				});
		}

		// Check for app updates on startup if enabled (non-blocking)
		const checkAppUpdatesOnStartup = getSetting('checkAppUpdatesOnStartup');
		if (checkAppUpdatesOnStartup !== false) { // Default to true if not set
			CheckForUpdates()
				.then((info) => {
					if (info) {
						updateAvailable.set(info);
					}
				})
				.catch((err) => {
					console.error('Background app update check failed:', err);
				});
		}

		// Cleanup function
		return () => {
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
