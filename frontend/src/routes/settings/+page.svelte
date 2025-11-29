<script lang="ts">
	import { onMount } from 'svelte';
	import { m } from '$lib/paraglide/messages.js';
	import { setLocale, getLocale, locales } from '$lib/paraglide/runtime.js';
	import {
		WindowSetSize,
		EventsOn,
		EventsOff,
		LogPrint,
		BrowserOpenURL
	} from '../../../wailsjs/runtime/runtime';
	import {
		CheckDefinitionsExist,
		CheckForDefsUpdates,
		UpdateDefinitions,
		GetDefinitionsPath,
		OpenAppDir,
		GetOSName,
		GetConfig,
		SaveSetting
	} from '../../../wailsjs/go/main/App';
	import {
		Download,
		RefreshCw,
		FolderOpen,
		CircleCheck,
		CircleAlert,
		Info,
		Languages,
		Monitor,
		Search,
		List,
		Funnel,
		RotateCcw,
		TriangleAlert,
		Smile,
		Shredder,
		Bug
	} from '@lucide/svelte';
	import { searchEngines } from '$lib/config/searchEngines';

	// Settings storage with defaults
	type Settings = {
		theme: 'light' | 'dark' | 'triduilight' | 'triduidark' | 'auto';
		searchEngine: string;
		maxVisibleMatches: number;
		confidenceThreshold: number;
		maxTotalResults: number;
		autoUpdateDefinitions: boolean;
		checkAppUpdatesOnStartup: boolean;
		startPinned: boolean;
		languageManuallySet: boolean;
	};

	const defaultSettings: Settings = {
		theme: 'auto',
		searchEngine: 'Google',
		maxVisibleMatches: 5,
		confidenceThreshold: 10.0,
		maxTotalResults: 50,
		autoUpdateDefinitions: true,
		checkAppUpdatesOnStartup: true,
		startPinned: false,
		languageManuallySet: false
	};

	if (typeof localStorage === 'undefined' || localStorage === null) {
		window.location.href = '/';
	}

	// LocalStorage helper functions
	function getSetting<K extends keyof Settings>(key: K): Settings[K] {
		const settingsData = localStorage.getItem('_trid_settings_');
		if (!settingsData) return defaultSettings[key];

		const settings = JSON.parse(settingsData);
		const value = settings[key];

		// If value doesn't exist in settings, return default
		if (value === undefined || value === null) return defaultSettings[key];

		// JSON.parse already converts types correctly, just return the value
		return value as Settings[K];
	}

	function setSetting<K extends keyof Settings>(key: K, value: Settings[K]): void {
		const settings = JSON.parse(localStorage.getItem('_trid_settings_') || '{}');
		settings[key] = value;
		localStorage.setItem('_trid_settings_', JSON.stringify(settings));
		try {
			SaveSetting(String(key), value);
		} catch (err) {
			console.debug('SaveSetting call failed:', err);
		}
	}

	function resetAllSettings(): void {
		localStorage.removeItem('_trid_settings_');
	}

	// Component state
	let definitionsExist = false;
	let definitionsPath = '';

	// Typed shape for update information returned by CheckForDefsUpdates
	type UpdateInfo = {
		defsCount: number;
		lastUpdated?: string;
		currentMD5?: string;
		isUpToDate: boolean;
		error?: string;
	};

	let updateInfo: UpdateInfo | null = null;
	let isCheckingUpdates = false;
	let isUpdating = false;
	let updateProgress = 0;
	let updateMessage = '';
	let updateError = '';
	let currentLocale = getLocale();
	let currentTheme = getSetting('theme');
	let currentSearchEngine = getSetting('searchEngine');
	let maxVisibleMatches = getSetting('maxVisibleMatches');
	let confidenceThreshold = getSetting('confidenceThreshold');
	let maxTotalResults = getSetting('maxTotalResults');
	let autoUpdateDefinitions = getSetting('autoUpdateDefinitions');
	let checkAppUpdatesOnStartup = getSetting('checkAppUpdatesOnStartup');
	let startPinned = getSetting('startPinned');

	// Dynamically get language display names based on available locales
	const availableLanguages = locales.map((locale) => {
		try {
			const displayNames = new Intl.DisplayNames([locale], { type: 'language' });
			return { code: locale, name: displayNames.of(locale) || locale };
		} catch {
			return { code: locale, name: locale.toUpperCase() };
		}
	});

	// Generic change handler for settings
	function handleSettingChange<K extends keyof Settings>(
		key: K,
		value: Settings[K],
		callback?: () => void
	) {
		setSetting(key, value);
		if (callback) callback();
	}

	function handleLanguageChange(event: Event) {
		const newLocale = (event.target as HTMLSelectElement).value as typeof locales[number];
		if (locales.includes(newLocale)) {
			setSetting('languageManuallySet', true);
			setLocale(newLocale);
			currentLocale = newLocale;
			window.location.href = '/';
		}
	}

	function applyTheme(theme: 'light' | 'dark' | 'triduilight' | 'triduidark' | 'auto') {
		if (typeof window === 'undefined') return;

		if (theme === 'auto') {
			const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
			document.documentElement.setAttribute(
				'data-theme',
				prefersDark ? 'triduidark' : 'triduilight'
			);
		} else {
			document.documentElement.setAttribute('data-theme', theme);
		}

		LogPrint('Applied theme: ' + theme);
		setSetting('theme', theme);
	}

	function handleThemeChange(event: Event) {
		const newTheme = (event.target as HTMLSelectElement).value as Settings['theme'];
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
		updateMessage = m['settings.starting_download']();

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

	function resetSettings() {
		// Reset to defaults
		currentTheme = defaultSettings.theme;
		currentSearchEngine = defaultSettings.searchEngine;
		maxVisibleMatches = defaultSettings.maxVisibleMatches;
		confidenceThreshold = defaultSettings.confidenceThreshold;
		maxTotalResults = defaultSettings.maxTotalResults;
		autoUpdateDefinitions = defaultSettings.autoUpdateDefinitions;
		checkAppUpdatesOnStartup = defaultSettings.checkAppUpdatesOnStartup;
		startPinned = defaultSettings.startPinned;

		// Clear localStorage
		resetAllSettings();

		// Apply theme
		applyTheme(defaultSettings.theme);

		// Reload to reset language to browser default
		window.location.href = '/';
	}

	onMount(() => {
		// Apply saved theme
		applyTheme(currentTheme);

		// Listen for system theme changes when in auto mode
		const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
		const handleSystemThemeChange = () => {
			if (currentTheme === 'auto') applyTheme('auto');
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
				if (definitionsExist && checkAppUpdatesOnStartup) {
					return checkForUpdates();
				}
			})
			.then(() => {
				// Auto-update definitions if enabled and update is available
				if (autoUpdateDefinitions && updateInfo && !updateInfo.isUpToDate) {
					return downloadUpdates();
				}
			})
			.catch((err) => {
				console.error('Error during initialization:', err);
			});

		// Listen for update progress events
		type UpdateProgressEvent = {
			percentage?: number;
			message?: string;
			downloaded?: number;
			total?: number;
		};

		EventsOn('trid:update:progress', (data: UpdateProgressEvent) => {
			updateProgress = data.percentage || 0;
			updateMessage = data.message || '';
			// if message includes DOWNLOADING, replace it with localized string
			if (data.message && data.message.toUpperCase().includes('DOWNLOADING')) {
				updateMessage =
					m['settings.downloading']() +
					`: ${Math.floor((data.downloaded || 0) / 1024).toLocaleString()} KB / ${Math.floor((data.total || 0) / 1024).toLocaleString()} KB`;
			}
		});

		EventsOn('trid:update:complete', () => {
			updateMessage = m['settings.update_complete']();
		});

		// Handle clicks on links with data-browserOpen attribute
		const handleLinkClick = (event: MouseEvent) => {
			const target = event.target as HTMLElement;
			const anchor = target.closest('a[data-browserOpen="true"]');
			if (anchor && anchor instanceof HTMLAnchorElement) {
				event.preventDefault();
				BrowserOpenURL(anchor.href);
			}
		};

		document.addEventListener('click', handleLinkClick);

		return () => {
			EventsOff('trid:update:progress');
			EventsOff('trid:update:complete');
			const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
			mediaQuery.removeEventListener('change', handleSystemThemeChange);
			document.removeEventListener('click', handleLinkClick);
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
				<div class="divider mt-0">{m['settings.trid_definitions']()}</div>
				<div class="card bg-base-100 shadow-sm">
					<div class="card-body p-5 space-y-4">
						<!-- Definitions Status -->
						<div class="alert {definitionsExist ? 'alert-info' : 'alert-warning'} alert-soft">
							{#if definitionsExist}
								<Info class="h-5 w-5" />
								<div class="flex-1">
									<h3 class="font-bold">{m['settings.definitions_installed']()}</h3>
									<div class="text-sm text-warp max-w-md wrap-anywhere">
										{#if updateInfo}
											<div class="mt-1 space-y-1">
												<div>
													{m['settings.definitions']()}:
													<span class="select-text"
														>{updateInfo.defsCount.toLocaleString()} file types</span
													>
												</div>
												{#if updateInfo.lastUpdated}
													<div>
														{m['settings.last_updated']()}:
														<span class="select-text">{updateInfo.lastUpdated}</span>
													</div>
												{/if}
												{#if updateInfo.currentMD5 && updateInfo.currentMD5 !== 'none'}
													<div class="text-xs opacity-70">
														MD5: <span class="select-all">{updateInfo.currentMD5}</span>
													</div>
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
							<div
								class={'alert ' +
									(updateInfo.isUpToDate ? 'alert-success' : 'alert-warning') +
									' alert-soft'}
							>
								{#if updateInfo.isUpToDate}
									<CircleCheck class="h-5 w-5" />
								{:else}
									<CircleAlert class="h-5 w-5" />
								{/if}
								<div class="flex-1">
									<div class="flex items-center justify-between">
										<span class="font-semibold">
											{updateInfo.isUpToDate
												? m['settings.up_to_date']()
												: m['settings.update_available']()}
										</span>
									</div>
									{#if updateInfo.error}
										<div class="text-xs text-error mt-1">{updateInfo.error}</div>
									{/if}
								</div>
							</div>
						{/if}

						<!-- Error Display -->
						{#if updateError}
							<div class="alert alert-error alert-soft">
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
								aria-label={m['about.check_for_updates']()}
							>
								{#if isCheckingUpdates}
									<RefreshCw class="h-4 w-4 animate-spin" />
									<span>{m['about.checking_for_updates']()}</span>
								{:else}
									<RefreshCw class="h-4 w-4" />
									<span>{m['about.check_for_updates']()}</span>
								{/if}
							</button>

							<button
								class="btn"
								on:click={(e: MouseEvent) => {
									const btn = e.currentTarget as HTMLButtonElement | null;
									if (btn) btn.disabled = true;
									try {
										openAppDirectory();
									} catch (error) {
										console.error('Failed to open app directory:', error);
									} finally {
										// wait 2000ms before re-enabling
										setTimeout(() => {
											if (btn) btn.disabled = false;
										}, 3500);
									}
								}}
								aria-label={m['settings.open_app_directory']()}
							>
								<FolderOpen class="h-4 w-4" />
								{m['settings.open_app_directory']()}
							</button>

							{#if !definitionsExist || (updateInfo && !updateInfo.isUpToDate)}
								<button
									class="btn btn-success"
									on:click={downloadUpdates}
									disabled={isUpdating}
									aria-label={isUpdating
										? m['settings.updating']()
										: definitionsExist
											? m['settings.update_definitions']()
											: m['settings.download_definitions']()}
								>
									{#if isUpdating}
										<Download class="h-4 w-4 animate-bounce" />
										{m['settings.updating']()}
									{:else}
										<Download class="h-4 w-4" />
										{definitionsExist
											? m['settings.update_definitions']()
											: m['settings.download_definitions']()}
									{/if}
								</button>
							{/if}
						</div>

						<!-- Info Box -->
						<div class="text-xs opacity-70 mt-2 p-3 bg-base-300 rounded-lg">
							<p class="mb-1">
								<strong>{m['settings.note']()}</strong>
								{m['settings.definitions_info']()}
							</p>
							<p>
								{m['settings.downloaded_from']()}:
								<a
									href="http://mark0.net/soft-trid-e.html"
									data-browserOpen="true"
									class="link"
									aria-label="mark0.net">mark0.net</a
								>
							</p>
						</div>
					</div>
				</div>
			</div>

			<!-- Appearance Section -->
			<div class="divider mt-8">{m['settings.language']()} & {m['settings.theme']()}</div>
			<div class="grid md:grid-cols-2 gap-6">
				<!-- Language Card -->
				<div class="card bg-base-100 shadow-sm">
					<div class="card-body p-5">
						<h3 class="card-title text-base flex items-center gap-2 mb-3">
							<Languages class="h-5 w-5" />
							{m['settings.language']()}
						</h3>
						<div class="form-control">
							<label class="label" for="language-select">
								<span class="label-text text-sm">
									{m['settings.language_description']()}
								</span>
							</label>
							<select
								id="language-select"
								class="select select-bordered w-full"
								value={currentLocale}
								on:change={handleLanguageChange}
								aria-label={m['settings.language']()}
							>
								{#each availableLanguages as { code, name } (code)}
									<option value={code}>{name.charAt(0).toUpperCase() + name.slice(1)}</option>
								{/each}
							</select>
							{#if currentLocale !== 'en'}
								<div class="mt-2">
									<span class="label-text-alt text-xs opacity-70 italic">
										{m['settings.language_disclaimer']()}
									</span>
								</div>
							{/if}
							<div class="mt-2">
								<span class="label-text-alt text-xs opacity-70">
									{@html m['settings.language_help_translation']()}
								</span>
							</div>
						</div>
					</div>
				</div>

				<!-- Theme Card -->
				<div class="card bg-base-100 shadow-sm">
					<div class="card-body p-5">
						<h3 class="card-title text-base flex items-center gap-2 mb-3">
							<Monitor class="h-5 w-5" />
							{m['settings.theme']()}
						</h3>
						<div class="form-control">
							<label class="label" for="theme-select">
								<span class="label-text text-sm text-wrap max-w-md">
									{m['settings.theme_description']()}
								</span>
							</label>
							<select
								id="theme-select"
								class="select select-bordered w-full"
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

								<option value="triduilight"> TridUI Light </option>
								<option value="triduidark"> TridUI Dark </option>
							</select>
							<div class="mt-2">
								<span class="label-text-alt text-xs opacity-70 text-wrap max-w-md">
									{m['settings.theme_auto_description']()}
								</span>
							</div>
						</div>
						<div class="divider my-2"></div>
						<div class="form-control">
							<label class="label cursor-pointer justify-start gap-3">
								<input
									type="checkbox"
									class="checkbox checkbox-primary"
									bind:checked={startPinned}
									on:change={() => handleSettingChange('startPinned', startPinned)}
								/>
								<div>
									<span class="label-text font-medium text-wrap max-w-md"
										>{m['settings.start_pinned']()}</span
									>
									<p class="label-text-alt text-xs opacity-70 mt-1 text-wrap max-w-md">
										{m['settings.start_pinned_description']()}
									</p>
								</div>
							</label>
						</div>
					</div>
				</div>
			</div>

			<!-- Search & Display Section -->
			<div class="divider mt-8">
				{m['settings.search_engine']()} & {m['settings.scan_results']()}
			</div>
			<div class="card bg-base-100 shadow-sm">
				<div class="card-body p-5">
					<!-- Search Engine -->
					<div>
						<h3 class="card-title text-base flex items-center gap-2 mb-3">
							<Search class="h-5 w-5" />
							{m['settings.search_engine']()}
						</h3>
						<div class="form-control">
							<label class="label" for="search-engine-select">
								<span class="label-text text-sm text-wrap max-w-md">
									{m['settings.search_engine_description']()}
								</span>
							</label>
							<select
								id="search-engine-select"
								class="select select-bordered w-full"
								bind:value={currentSearchEngine}
								on:change={() => handleSettingChange('searchEngine', currentSearchEngine)}
								aria-label={m['settings.search_engine']()}
							>
								{#each searchEngines as engine (engine.name)}
									<option value={engine.name}>{engine.name}</option>
								{/each}
							</select>
						</div>
					</div>

					<div class="divider my-3"></div>

					<!-- Scan Results Display -->
					<div>
						<h3 class="card-title text-base flex items-center gap-2 mb-4">
							<List class="h-5 w-5" />
							{m['settings.scan_results']()}
						</h3>

						<div class="space-y-5">
							<!-- Max Visible Matches -->
							<div class="form-control">
								<label class="label" for="max-results-input">
									<span class="label-text text-sm font-medium text-wrap max-w-md">
										{m['settings.max_visible_matches']()}
									</span>
								</label>
								<input
									id="max-results-input"
									type="number"
									min="1"
									max="50"
									class="input input-bordered w-full"
									bind:value={maxVisibleMatches}
									on:change={() =>
										maxVisibleMatches >= 1 &&
										maxVisibleMatches <= 50 &&
										handleSettingChange('maxVisibleMatches', maxVisibleMatches)}
									aria-label={m['settings.max_visible_matches']()}
								/>
								<div class="label">
									<span class="label-text-alt text-xs opacity-70 text-wrap max-w-md">
										{m['settings.max_visible_matches_description']()}
									</span>
								</div>
							</div>

							<!-- Confidence Threshold -->
							<div class="form-control">
								<label class="label" for="threshold-input">
									<span
										class="label-text text-sm font-medium flex items-center gap-2 text-wrap max-w-md"
									>
										<Funnel class="h-4 w-4" />
										{m['settings.confidence_threshold']()}
									</span>
								</label>
								<div class="flex items-center gap-3">
									<input
										id="threshold-input"
										type="range"
										min="0"
										max="100"
										step="0.5"
										class="range range-primary flex-1"
										bind:value={confidenceThreshold}
										on:change={() =>
											handleSettingChange('confidenceThreshold', confidenceThreshold)}
										aria-label={m['settings.confidence_threshold']()}
									/>
									<span class="badge badge-neutral font-mono min-w-14 justify-center">
										{confidenceThreshold.toFixed(1)}%
									</span>
								</div>
								<div class="label">
									<span class="label-text-alt text-xs opacity-70 text-wrap max-w-md">
										{m['settings.confidence_threshold_description']()}
									</span>
								</div>
							</div>

							<!-- Max Total Results -->
							<div class="form-control">
								<label class="label" for="max-total-results-input">
									<span class="label-text text-sm font-medium text-wrap max-w-md">
										{m['settings.max_total_results']()}
									</span>
								</label>
								<input
									id="max-total-results-input"
									type="number"
									min="0"
									max="10000"
									class="input input-bordered w-full"
									bind:value={maxTotalResults}
									on:change={() =>
										maxTotalResults >= 0 &&
										maxTotalResults <= 10000 &&
										handleSettingChange('maxTotalResults', maxTotalResults)}
									aria-label={m['settings.max_total_results']()}
									aria-describedby="max-total-results-desc"
								/>
								<div class="label">
									<span
										id="max-total-results-desc"
										class="label-text-alt text-xs opacity-70 text-wrap max-w-md"
									>
										{m['settings.max_total_results_description']()}
									</span>
								</div>
								<div class="alert alert-warning alert-soft py-2 px-3 mt-2 text-wrap max-w-md">
									<TriangleAlert class="h-4 w-4" />
									<span class="text-xs">{m['settings.max_total_results_warning']()}</span>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>

			<!-- Updates Section -->
			<div class="divider mt-8">{m['settings.updates']()}</div>
			<div class="card bg-base-100 shadow-sm">
				<div class="card-body p-5 space-y-4">
					<div class="form-control">
						<label class="label cursor-pointer justify-start gap-3">
							<input
								type="checkbox"
								class="checkbox checkbox-primary"
								bind:checked={checkAppUpdatesOnStartup}
								on:change={() =>
									handleSettingChange('checkAppUpdatesOnStartup', checkAppUpdatesOnStartup)}
							/>
							<div>
								<span class="label-text font-medium text-wrap max-w-md"
									>{m['settings.check_app_updates']()}</span
								>
								<p class="label-text-alt text-xs opacity-70 mt-1 text-wrap max-w-md">
									{m['settings.check_app_updates_description']()}
								</p>
							</div>
						</label>
					</div>

					<!-- Auto-Update Definitions -->
					<div>
						<div class="form-control">
							<label class="label cursor-pointer justify-start gap-3">
								<input
									type="checkbox"
									class="checkbox checkbox-primary"
									bind:checked={autoUpdateDefinitions}
									on:change={() =>
										handleSettingChange('autoUpdateDefinitions', autoUpdateDefinitions)}
								/>
								<div>
									<span class="label-text font-medium text-wrap max-w-md"
										>{m['settings.auto_update_definitions']()}</span
									>
									<p class="label-text-alt text-xs opacity-70 mt-1 text-wrap max-w-md">
										{m['settings.auto_update_definitions_description']()}
									</p>
								</div>
							</label>
						</div>
					</div>
				</div>
			</div>

			<!-- Debug Section -->
			<details class="collapse border-base-300 bg-base-100 border mt-6">
				<summary class="collapse-title text-sm opacity-70"
					><Bug class="inline-block h-4 w-4" /> {m['settings.debug_tools']()}</summary
				>
				<div class="collapse-content text-sm space-y-2">
					<button class="btn btn-sm btn-primary" on:click={() => WindowSetSize(500, 400)}>
						Reset window size
					</button>
					<br />
					<button
						class="btn btn-sm btn-secondary"
						on:click={(e) => {
							const btn = e.currentTarget as HTMLButtonElement;
							navigator.clipboard
								.writeText(localStorage.getItem('_trid_settings_') || '{}')
								.then(() => {
									btn.innerText = 'Settings copied to clipboard!';
								})
								.catch((err) => {
									btn.innerText = 'Failed to copy settings: ' + err;
								})
								.finally(() => {
									setTimeout(() => {
										btn.innerText = 'Copy settings to clipboard';
									}, 3000);
								});
						}}
					>
						Copy settings to clipboard
					</button>

					<button
						class="btn btn-sm btn-ghost"
						on:click={async (e) => {
							const btn = e.currentTarget as HTMLButtonElement;
							btn.disabled = true;
							const orig = btn.innerText;
							btn.innerText = 'Reloading...';
							try {
								const cfg = await GetConfig();
								if (!cfg) throw new Error('empty config');
								const parsed = JSON.parse(cfg || '{}') as Record<string, unknown>;
								localStorage.setItem('_trid_settings_', JSON.stringify(parsed));
								// update component state from loaded settings (with typed casts)
								currentTheme = (parsed.theme as Settings['theme']) || defaultSettings.theme;
								currentSearchEngine = (parsed.searchEngine as string) || defaultSettings.searchEngine;
								maxVisibleMatches = (parsed.maxVisibleMatches as number) ?? defaultSettings.maxVisibleMatches;
								confidenceThreshold = (parsed.confidenceThreshold as number) ?? defaultSettings.confidenceThreshold;
								maxTotalResults = (parsed.maxTotalResults as number) ?? defaultSettings.maxTotalResults;
								autoUpdateDefinitions = (parsed.autoUpdateDefinitions as boolean) ?? defaultSettings.autoUpdateDefinitions;
								checkAppUpdatesOnStartup = (parsed.checkAppUpdatesOnStartup as boolean) ?? defaultSettings.checkAppUpdatesOnStartup;
								startPinned = (parsed.startPinned as boolean) ?? defaultSettings.startPinned;
								// apply theme and language if present
								applyTheme(currentTheme);
								if (typeof parsed.language === 'string') {
									const lang = parsed.language as typeof locales[number];
									if (locales.includes(lang)) {
										setLocale(lang);
									}
								}
								btn.innerText = 'Reloaded';
							} catch (err) {
								console.error('Failed to reload settings from file:', err);
								btn.innerText = 'Reload failed';
							} finally {
								setTimeout(() => {
									btn.disabled = false;
									btn.innerText = orig;
								}, 1500);
							}
						}}
					>
						Reload settings from file
					</button>
					<br />
					<button
						class="btn btn-sm btn-secondary text-wrap wrap-anywhere"
						on:click={async (e) => {
							const btn = e.currentTarget as HTMLButtonElement;
							btn.disabled = true;
							// hide the first span and show the second span
							btn.children[0].classList.add('hidden');
							btn.children[1].classList.remove('hidden');
							await new Promise((resolve) => setTimeout(resolve, 4000));
							// remove hidden from children 1 child and add hidden to children 0 child
							btn.children[1].children[0].classList.remove('hidden');
							await new Promise((resolve) => setTimeout(resolve, 2000));
							btn.children[1].classList.add('hidden');
							btn.children[1].children[0].classList.add('hidden');
							btn.children[0].classList.remove('hidden');
							btn.disabled = false;
						}}
					>
						<span
							><Shredder class="inline-block h-4 w-4" />Delete {#await GetOSName() then os}
								{os === 'windows' ? 'C:/Windows/System32' : os === 'macos' ? '/System' : '/root'}
							{/await} directory (for testing)</span
						>
						<span class="hidden"
							>Elevating process... <span class="hidden"
								>Just kidding! <Smile class="inline-block h-4 w-4" /></span
							></span
						>
					</button>
					<br />
					{#if definitionsPath}
						<div class="text-xs opacity-70 break-all">
							<strong>{m['settings.located_at']()}:</strong>
							<span class="select-text text-wrap max-w-md wrap-anywhere">{definitionsPath}</span>
						</div>
					{/if}
				</div>
			</details>

			<!-- Reset Settings Section -->
			<div class="divider mt-2"></div>
			<div class="card bg-base-100 shadow-sm border-warning border">
				<div class="card-body p-5">
					<h3 class="card-title text-base flex items-center gap-2 mb-3">
						<RotateCcw class="h-5 w-5" />
						{m['settings.reset_settings']()}
					</h3>
					<p class="text-sm opacity-70 mb-4">
						{m['settings.reset_settings_description']()}
					</p>
					<button class="btn btn-warning" on:click={resetSettings}>
						<RotateCcw class="h-4 w-4" />
						{m['settings.reset_to_defaults']()}
					</button>
				</div>
			</div>
		</div>
	</div>
</div>
