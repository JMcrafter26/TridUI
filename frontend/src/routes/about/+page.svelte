<script lang="ts">
	import { onMount } from 'svelte';
	import { m } from '$lib/paraglide/messages.js';
	import { updateAvailable } from '$lib/stores/updateStore';
	import logo from '$lib/images/app-icon.png';
	import { Github, Download, RefreshCw, CircleCheck, CircleAlert, X, Rocket, ChevronDown, ChevronUp } from '@lucide/svelte';
	import { GetVersion, CheckForUpdates } from '../../../wailsjs/go/main/App';
	import { BrowserOpenURL } from '../../../wailsjs/runtime/runtime';
	import type { main } from '../../../wailsjs/go/models';
	import snarkdown from 'snarkdown';

	let version = '1.0.0';
	let updateInfo: main.UpdateInfo | null = null;
	let isChecking = false;
	let errorTitle = '';
	let errorMessage = '';
	let showReleaseNotes = false;
	let showSuccessAlert = false;

	onMount(async () => {
		try {
			version = await GetVersion();
			// Check if we already have update info from startup check
			if ($updateAvailable) {
				updateInfo = $updateAvailable;
				// Scroll to the update alert after a short delay if there's an update
				if (updateInfo.updateAvailable) {
					setTimeout(() => {
						document.getElementById('update-available')?.scrollIntoView({
							behavior: 'smooth',
							block: 'center'
						});
					}, 100);
				}
			}
		} catch (err) {
			console.error('Failed to get version:', err);
		}
	});

	onMount(() => {
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
			document.removeEventListener('click', handleLinkClick);
		};
	});

	async function checkForUpdates() {
		isChecking = true;
		errorTitle = '';
		errorMessage = '';
		updateInfo = null;
		showReleaseNotes = false;
		showSuccessAlert = true;

		try {
			updateInfo = await CheckForUpdates();
			// Update the store for the indicator
			if (updateInfo && updateInfo.updateAvailable) {
				updateAvailable.set(updateInfo);
			} else {
				updateAvailable.set(null);
			}
		} catch (err) {
			console.error('Update check failed - full error:', err);

			errorTitle = m['about.update_check_failed']();

			if (err instanceof Error) {
				errorMessage = err.message;
			} else if (typeof err === 'string') {
				errorMessage = err;
			} else if (err && typeof err === 'object') {
				// Handle Wails error object
				errorMessage = (err as any).message || JSON.stringify(err);
			} else {
				errorMessage = 'An unknown error occurred';
			}
		} finally {
			isChecking = false;
		}
	}

	function openReleaseURL() {
		if (updateInfo?.releaseUrl) {
			window.open(updateInfo.releaseUrl, '_blank');
		}
	}

	function formatDate(dateStr: string): string {
		if (!dateStr) return '';
		try {
			return new Date(dateStr).toLocaleDateString(undefined, {
				year: 'numeric',
				month: 'long',
				day: 'numeric'
			});
		} catch {
			return dateStr;
		}
	}

	function formatReleaseNotes(markdown: string): string {
		if (!markdown) return '';
		// Convert markdown to HTML
		let html = snarkdown(markdown);
		// Add target="_blank" to all links and rel="noopener noreferrer" and BrowserOpenURL(url) on:click
		// Replace anchors with buttons that open the URL in a new window
		html = html.replace(/<a\s+href="([^"]*)"(.*?)>/gi, (_match, href, attrs) => {
			const escaped = href.replace(/'/g, "\\'");
			return `<a class="link cursor-pointer" href="${escaped}" data-browserOpen="true" ${attrs}>`;
		});
		// Close buttons for closing anchor tags
		// html = html.replace(/<\/a>/gi, '</button>');
		return html;
	}

	function toggleReleaseNotes() {
		showReleaseNotes = !showReleaseNotes;
	}
</script>

<svelte:head>
	<title>{m['about.about']()}</title>
	<meta name="description" content={m['about.description']()} />
</svelte:head>

<div class="max-w-4xl mx-auto px-4 py-4 h-full">
	<div class="card bg-base-200 shadow-lg h-full overflow-auto">
		<div class="card-body p-6 md:p-10 flex flex-col md:flex-row items-center gap-6 mb-4">
			<div class="shrink-0">
				<img src={logo} alt="App icon" class="w-32 h-32 object-cover transition-all hover:scale-110 hover:rotate-3" />
			</div>

			<div class="flex-1 min-w-0">
				<h1 class="text-2xl md:text-3xl font-semibold leading-tight">TrID UI</h1>
				<div class="text-sm text-muted mt-1">
					{m['about.version']()} <span class="font-medium select-all">{version}</span>
				</div>

				<p class="mt-4 text-sm text-base-content/80">{m['about.description']()}</p>

				<div class="mt-6 flex flex-wrap gap-3 items-center">
					<button
						class="btn btn-primary btn-sm"
						on:click={checkForUpdates}
						disabled={isChecking}
						aria-label={m['about.check_for_updates']()}
					>
						{#if isChecking}
							<RefreshCw class="h-4 w-4 animate-spin" />
							{m['about.checking_for_updates']()}
						{:else}
							<RefreshCw class="h-4 w-4" />
							{m['about.check_for_updates']()}
						{/if}
					</button>

					<button
						class="inline-flex items-center gap-2 btn btn-sm"
						on:click={() => BrowserOpenURL('https://github.com/JMcrafter26/TridUI')}
						aria-label={m['about.view_on_github']()}
					>
						<span class="w-4 h-4 opacity-80">
							<Github />
						</span>
						<span>{m['about.view_on_github']()}</span>
					</button>
				</div>
			</div>

			<!-- Update Check Results -->
			{#if errorTitle}
				<div
					class="alert alert-soft alert-error mt-4 fixed bottom-3 md:bottom-6 max-w-md flex items-start gap-4 mr-20 ml-4"
				>
					<CircleAlert class="h-5 w-5 my-auto" />
					<div class="flex-1 select-text">
						<h3 class="font-bold">{errorTitle}</h3>
						{#if errorMessage}
							<div class="text-xs mt-1">{errorMessage}</div>
						{/if}
					</div>
					<!-- close X -->
					<button
						class="btn btn-ghost btn-sm hover:text-error hover:bg-error/10 rounded-full p-1 my-auto"
						on:click={() => {
							errorTitle = '';
							errorMessage = '';
						}}
					>
						<X />
					</button>
				</div>
			{/if}
			{#if updateInfo}
				{#if updateInfo.updateAvailable}
					<div class="alert alert-soft alert-info mt-4" id="update-available">
						<Rocket class="h-6 w-6" />
						<div class="flex-1">
							<h3 class="font-bold">{m['about.update_available']()}</h3>
							<div class="text-sm mt-1">
								<p>
									{m['about.new_version_available']({ latestVersion: updateInfo.latestVersion })}
									{#if updateInfo.publishedAt}
										({m['about.released']()} {formatDate(updateInfo.publishedAt)})
									{/if}
								</p>
							</div>
						</div>

						<button
							class="btn btn-sm btn-primary"
							on:click={openReleaseURL}
							aria-label={m['about.download']()}
						>
							<Download class="h-4 w-4" />
							{m['about.download']()}
						</button>
					</div>

					{#if updateInfo.releaseNotes}
						<div class="mt-1 p-4 bg-base-300 rounded-md max-h-80 overflow-auto w-full">
							<h4 class="font-medium mb-2">{m['about.release_notes']()}:</h4>
							<div
								class="text-sm space-y-2 [&_h1]:text-xl [&_h1]:font-bold [&_h2]:text-lg [&_h2]:font-semibold [&_h3]:text-base [&_h3]:font-medium [&_p]:mb-2 [&_ul]:list-disc [&_ul]:ml-4 [&_ul]:space-y-1 [&_ol]:list-decimal [&_ol]:ml-4 [&_ol]:space-y-1 [&_a]:text-primary [&_a]:underline [&_a]:hover:opacity-80 [&_strong]:font-semibold [&_em]:italic [&_code]:bg-base-100 [&_code]:px-1 [&_code]:rounded"
							>
								{@html formatReleaseNotes(updateInfo.releaseNotes)}
							</div>
							<a
								href={updateInfo.releaseUrl}
								target="_blank"
								rel="noopener noreferrer"
								class="btn btn-sm btn-secondary mt-4 inline-flex items-center gap-2"
							>
								<Github class="h-4 w-4" />
								{m['about.view_on_github']()}
							</a>
						</div>
					{/if}
				{:else}
					{#if showSuccessAlert}
						<div class="alert alert-soft alert-success mt-4 fixed bottom-3 md:bottom-6 max-w-md">
							<CircleCheck class="h-5 w-5 my-auto" />
							<span>{m['about.up_to_date']()}</span>
							<button
								class="btn btn-ghost btn-sm rounded-full p-1 my-auto hover:bg-success/10"
								on:click={() => (showSuccessAlert = false)}
							>
								<X />
							</button>
						</div>
					{/if}
				{/if}

			{#if updateInfo && !updateInfo.updateAvailable && updateInfo.releaseNotes}
				<details class="mt-4 w-full group" open={showReleaseNotes}>
							<summary
								class="cursor-pointer select-none p-3 bg-base-300 rounded-lg hover:bg-base-300/80 transition-colors flex items-center justify-between sticky top-0

								group-open:rounded-b-none"
								on:click|preventDefault={toggleReleaseNotes}
							>
								<span class="font-medium flex items-center gap-2">
									{m['about.current_release_notes']()}
									{#if updateInfo.publishedAt && !showReleaseNotes}
										<span class="text-xs text-base-content/70 font-normal">
											({formatDate(updateInfo.publishedAt)})
										</span>
								{/if}
							</span>
							{#if showReleaseNotes}
								<ChevronUp class="h-4 w-4" />
							{:else}
								<ChevronDown class="h-4 w-4" />
							{/if}
							</summary>
					{#if showReleaseNotes}
						<div class="p-4 bg-base-300/50 rounded-b-md max-h-80 overflow-auto">
									<div class="text-sm mb-2">
										<span class="font-medium">Version {updateInfo.currentVersion}</span>
										{#if updateInfo.publishedAt}
											<span class="text-base-content/70"> • {formatDate(updateInfo.publishedAt)}</span>
										{/if}
									</div>
									<div
										class="text-sm space-y-2 [&_h1]:text-xl [&_h1]:font-bold [&_h2]:text-lg [&_h2]:font-semibold [&_h3]:text-base [&_h3]:font-medium [&_p]:mb-2 [&_ul]:list-disc [&_ul]:ml-4 [&_ul]:space-y-1 [&_ol]:list-decimal [&_ol]:ml-4 [&_ol]:space-y-1 [&_a]:text-primary [&_a]:underline [&_a]:hover:opacity-80 [&_strong]:font-semibold [&_em]:italic [&_code]:bg-base-100 [&_code]:px-1 [&_code]:rounded text-wrap wrap-anywhere"
									>
										{@html formatReleaseNotes(updateInfo.releaseNotes)}
								</div>
								{#if updateInfo.releaseUrl}
									<button
										on:click={() => BrowserOpenURL(updateInfo?.releaseUrl || '')}
										class="btn btn-sm btn-secondary mt-4 inline-flex items-center gap-2"
									>
										<Github class="h-4 w-4" />
										{m['about.view_on_github']()}
									</button>
									{/if}
								</div>
							{/if}
						</details>
					{/if}
				{/if}

			<div class="divider md:divider-horizontal"></div>

			<div class="w-full md:w-56">
				<h3 class="text-sm font-medium">{m['about.about_and_credits']()}</h3>
				<p class="text-xs text-base-content/70 mt-2">{@html m['about.made_by']()}</p>

				<div class="mt-4 text-xs text-base-content/70 mb-3">
					<div>{m['about.license']()}: GNU AGPL v3.0</div>
					<div class="mt-2">{m['about.built_with']()}</div>
				</div>

				<div class="mt-4 text-xs text-base-content/50 italic text-center mb-0">
					<p class="transition-all hover:rotate-x-180">{m['about.made_with_hearts_by']()}</p>
				</div>
			</div>
		</div>
	</div>
</div>
