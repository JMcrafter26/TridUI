<script lang="ts">
	import { onMount } from 'svelte';
	import { m } from '$lib/paraglide/messages.js';
	import { setLocale, getLocale, locales } from '$lib/paraglide/runtime.js';
	import { WindowSetSize, EventsOn, EventsOff } from '../../../wailsjs/runtime/runtime';
	import {
		CheckDefinitionsExist,
		CheckForDefsUpdates,
		UpdateDefinitions,
		GetDefinitionsPath,
		OpenAppDir
	} from '../../../wailsjs/go/main/App';
	import { Download, RefreshCw, FolderOpen, CircleCheck, CircleAlert, Info, Languages, Moon, Sun, Monitor } from '@lucide/svelte';

	let definitionsExist = false;
	let definitionsPath = '';
	let updateInfo: any | null = null;
	let isCheckingUpdates = false;
	let isUpdating = false;
	let updateProgress = 0;
	let updateMessage = '';
	let updateError = '';
	let currentLocale = getLocale();
	let currentTheme: 'light' | 'dark' | 'auto' = 'auto';

	// Dynamically get language display names based on available locales
	const availableLanguages = locales.map((locale) => {
		try {
			// Use Intl.DisplayNames to get native language names
			const displayNames = new Intl.DisplayNames([locale], { type: 'language' });
			return {
				code: locale,
				name: displayNames.of(locale) || locale
			};
		} catch {
			// Fallback if Intl.DisplayNames fails
			return {
				code: locale,
				name: locale.toUpperCase()
			};
		}
	});

	function handleLanguageChange(event: Event) {
		const target = event.target as HTMLSelectElement;
		const newLocale = target.value;
		if (locales.includes(newLocale as any)) {
			// Mark that user has manually selected a language
			localStorage.setItem('language-manually-set', 'true');
			setLocale(newLocale as any);
			currentLocale = newLocale as any;
			// goto /
			window.location.href = '/';
		}
	}

	function applyTheme(theme: 'light' | 'dark' | 'auto') {
		if (typeof window === 'undefined') return;
		
		if (theme === 'auto') {
			const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
			document.documentElement.setAttribute('data-theme', prefersDark ? 'dark' : 'light');
		} else {
			document.documentElement.setAttribute('data-theme', theme);
		}
		
		localStorage.setItem('theme', theme);
	}

	function handleThemeChange(event: Event) {
		const target = event.target as HTMLSelectElement;
		const newTheme = target.value as 'light' | 'dark' | 'auto';
		currentTheme = newTheme;
		applyTheme(newTheme);
	}

	async function checkForUpdates() {
		isCheckingUpdates = true;
		updateError = '';
		try {
			updateInfo = await CheckForDefsUpdates();
		} catch (err) {
			updateError = String(err);
			console.error('Failed to check for updates:', err);
		} finally {
			isCheckingUpdates = false;
		}
	}

	async function downloadUpdates() {
		isUpdating = true;
		updateError = '';
		updateProgress = 0;
		updateMessage = 'Starting download...';

		try {
			await UpdateDefinitions();
			// Refresh update info after successful update
			await checkForUpdates();
			definitionsExist = await CheckDefinitionsExist();
		} catch (err) {
			updateError = String(err);
			console.error('Failed to update definitions:', err);
		} finally {
			isUpdating = false;
		}
	}

	async function openAppDirectory() {
		try {
			await OpenAppDir();
		} catch (err) {
			console.error('Failed to open app directory:', err);
		}
	}

	onMount(() => {
		// Load saved theme
		const savedTheme = localStorage.getItem('theme') as 'light' | 'dark' | 'auto' | null;
		if (savedTheme) {
			currentTheme = savedTheme;
			applyTheme(savedTheme);
		} else {
			applyTheme('auto');
		}

		// Listen for system theme changes when in auto mode
		const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
		const handleSystemThemeChange = () => {
			if (currentTheme === 'auto') {
				applyTheme('auto');
			}
		};
		mediaQuery.addEventListener('change', handleSystemThemeChange);

		// Check if definitions exist
		CheckDefinitionsExist()
			.then((exists) => {
				definitionsExist = exists;
				return GetDefinitionsPath();
			})
			.then((path) => {
				definitionsPath = path;
				// Auto-check for updates if definitions exist
				if (definitionsExist) {
					return checkForUpdates();
				}
			})
			.catch((err) => {
				console.error('Error during initialization:', err);
			});

		// Listen for update progress events
		EventsOn('trid:update:progress', (data: any) => {
			updateProgress = data.percentage || 0;
			updateMessage = data.message || '';
		});

		EventsOn('trid:update:complete', () => {
			updateMessage = 'Update completed successfully!';
		});

		return () => {
			EventsOff('trid:update:progress');
			EventsOff('trid:update:complete');
			const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
			mediaQuery.removeEventListener('change', handleSystemThemeChange);
		};
	});
	</script>

<svelte:head>
	<title>{m['settings.settings']()}</title>
	<meta name="description" content={m['settings.description']()} />
</svelte:head>

<div class="max-w-4xl mx-auto px-4 py-4 h-full">
	<div class="card bg-base-200 shadow-lg h-full overflow-auto">
		<div class="card-body p-6 md:p-10">
			<h2 class="card-title text-xl mb-4">{m['settings.settings']()}</h2>

			<!-- TrID Definitions Section -->
			<div class="space-y-4">
				<div class="divider">{m['settings.trid_definitions']()}</div>

				<!-- Definitions Status -->
				<div class="alert {definitionsExist ? 'alert-success' : 'alert-warning'}">
					{#if definitionsExist}
						<CircleCheck class="h-5 w-5" />
						<div class="flex-1">
							<h3 class="font-bold">{m['settings.definitions_installed']()}</h3>
							<div class="text-sm">
								{#if updateInfo}
									<div class="mt-1 space-y-1">
										<div>{m['settings.definitions']()}: <span class="select-text">{updateInfo.defsCount.toLocaleString()}</span> file types</div>
										{#if updateInfo.lastUpdated}
											<div>{m['settings.last_updated']()}: <span class="select-text">{updateInfo.lastUpdated}</span></div>
										{/if}
										{#if updateInfo.currentMD5 && updateInfo.currentMD5 !== 'none'}
											<div class="text-xs opacity-70">MD5: <span class="select-all">{updateInfo.currentMD5}</span></div>
										{/if}
									</div>
								{:else}
									{m['settings.located_at']()}: {definitionsPath}
								{/if}
							</div>
						</div>
					{:else}
						<CircleAlert class="h-5 w-5" />
						<div class="flex-1">
							<h3 class="font-bold">{m['settings.definitions_not_found']()}</h3>
							<div class="text-sm">
								{m['settings.definitions_not_found_explanation']()}
							</div>
						</div>
					{/if}
				</div>

				<!-- Update Status -->
				{#if updateInfo && definitionsExist}
					<div class="card bg-base-300">
						<div class="card-body p-4">
							<div class="flex items-center justify-between">
								<div class="flex items-center gap-2">
									<Info class="h-5 w-5" />
									<span class="font-semibold">
										{updateInfo.isUpToDate ? m["settings.up_to_date"]() : m["settings.update_available"]()}
									</span>
								</div>
								{#if !updateInfo.isUpToDate}
									<span class="badge badge-primary">{m['settings.new_version']()}</span>
								{/if}
							</div>
							{#if updateInfo.error}
								<div class="text-xs text-error mt-2">{updateInfo.error}</div>
							{/if}
						</div>
					</div>
				{/if}

				<!-- Error Display -->
				{#if updateError}
					<div class="alert alert-error">
						<CircleAlert class="h-5 w-5" />
						<div>
							<h3 class="font-bold">{m['settings.error']()}</h3>
							<div class="text-sm">{updateError}</div>
						</div>
					</div>
				{/if}

				<!-- Update Progress -->
				{#if isUpdating}
					<div class="card bg-base-300">
						<div class="card-body p-4">
							<div class="flex items-center gap-3 mb-2">
								<Download class="h-5 w-5 animate-bounce" />
								<span class="font-semibold">{updateMessage}</span>
							</div>
							<progress
								class="progress progress-primary w-full"
								value={updateProgress}
								max="100"
							></progress>
							<div class="text-xs text-center mt-1">{updateProgress}%</div>
						</div>
					</div>
				{/if}

				<!-- Action Buttons -->
				<div class="flex flex-wrap gap-2">
					<button
						class="btn btn-primary"
						on:click={checkForUpdates}
						disabled={isCheckingUpdates || isUpdating}
						aria-label={m["about.check_for_updates"]()}
					>
						{#if isCheckingUpdates}
							<RefreshCw class="h-4 w-4 animate-spin" />
							<span>{m["about.checking_for_updates"]()}</span>
						{:else}
							<RefreshCw class="h-4 w-4" />
							<span>{m["about.check_for_updates"]()}</span>
						{/if}
					</button>

					{#if !definitionsExist || (updateInfo && !updateInfo.isUpToDate)}
						<button class="btn btn-success" on:click={downloadUpdates} disabled={isUpdating} aria-label="{isUpdating ? m["settings.updating"]() : definitionsExist ? m["settings.update_definitions"]() : m["settings.download_definitions"]()}">
							{#if isUpdating}
								<Download class="h-4 w-4 animate-bounce" />
								{m["settings.updating"]()}
							{:else}
								<Download class="h-4 w-4" />
								{definitionsExist ? m["settings.update_definitions"]() : m["settings.download_definitions"]()}
							{/if}
						</button>
					{/if}

					<button class="btn btn-ghost" on:click={openAppDirectory} aria-label={m["settings.open_app_directory"]()}>
						<FolderOpen class="h-4 w-4" />
						{m['settings.open_app_directory']()}
					</button>
				</div>

				<!-- Info Box -->
				<div class="text-xs opacity-70 mt-4 p-3 bg-base-300 rounded">
					<p class="mb-1">
						<strong>{m['settings.note']()}</strong> {m['settings.definitions_info']()}
					</p>
					<p>
						{m['settings.downloaded_from']()}:
						<a
							href="http://mark0.net/soft-trid-e.html"
							target="_blank"
							rel="noopener noreferrer"
							class="link"
							aria-label="mark0.net"
							>mark0.net</a
						>
					</p>
				</div>
			</div>

			<!-- Language Section -->
			<div class="space-y-4">
				<div class="divider">{m['settings.language']()}</div>

				<div class="form-control w-full max-w-xs">
					<label class="label" for="language-select">
						<span class="label-text flex items-center gap-2">
							<Languages class="h-4 w-4" />
							{m['settings.language_description']()}
						</span>
					</label>
					<select
						id="language-select"
						class="select select-bordered w-full max-w-xs"
						value={currentLocale}
						on:change={handleLanguageChange}
						aria-label={m['settings.language']()}
					>
						{#each availableLanguages as { code, name }}
						<!-- capitalize the first letter of the language name -->
							<option value={code}>{name.charAt(0).toUpperCase() + name.slice(1)}</option>
						{/each}
					</select>
					<!-- if language is not english, show disclaimer -->
					{#if currentLocale !== 'en'}
						<div class="mt-1">
							<span class="label-text-alt text-xs opacity-70 italic">
								{m['settings.language_disclaimer']()}
							</span>
						</div>
					{/if}

					<!-- help translate -->
					<div class="mt-1">
						<span class="label-text-alt text-xs opacity-70">
							{@html m['settings.language_help_translation']()} 
						</span>
					</div>
				</div>
			</div>

			<!-- Theme Section -->
			<div class="space-y-4">
				<div class="divider">{m['settings.theme']()}</div>

				<div class="form-control w-full max-w-xs">
					<label class="label" for="theme-select">
						<span class="label-text flex items-center gap-2">
							<Monitor class="h-4 w-4" />
							{m['settings.theme_description']()}
						</span>
					</label>
					<select
						id="theme-select"
						class="select select-bordered w-full max-w-xs"
						value={currentTheme}
						on:change={handleThemeChange}
					>
						<option value="light">
							{m['settings.theme_light']()}
						</option>
						<option value="dark">
							{m['settings.theme_dark']()}
						</option>
						<option value="auto">
							{m['settings.theme_auto']()}
						</option>
					</select>
					<div class="mt-1">
						<span class="label-text-alt text-xs opacity-70">
							{m['settings.theme_auto_description']()}
						</span>
					</div>
				</div>
			</div>

			<!-- Debug Section -->
			<details class="collapse border-base-300 border mt-6">
				<summary class="collapse-title text-sm opacity-70">{m['settings.debug_tools']()}</summary>
				<div class="collapse-content text-sm space-y-2">
					<button class="btn btn-sm btn-primary" on:click={() => WindowSetSize(500, 400)}>
						Reset window size
					</button>
					{#if definitionsPath}
						<div class="text-xs opacity-70 break-all">
							<strong>{m['settings.located_at']()}:</strong>
							<span class="select-text">{definitionsPath}</span>
						</div>
					{/if}
				</div>
			</details>
		</div>
	</div>
</div>
