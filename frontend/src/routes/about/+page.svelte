<script lang="ts">
	import { onMount } from 'svelte';
	import { m } from '$lib/paraglide/messages.js';
	import logo from '$lib/images/app-icon.png';
	import { Github, Download, RefreshCw, CircleCheck, CircleAlert, X } from '@lucide/svelte';
	import { GetVersion, CheckForUpdates } from '../../../wailsjs/go/main/App';
	import type { main } from '../../../wailsjs/go/models';

	let version = '1.0.0';
	let updateInfo: main.UpdateInfo | null = null;
	let isChecking = false;
	let errorTitle = '';
	let errorMessage = '';

	onMount(async () => {
		try {
			version = await GetVersion();
		} catch (err) {
			console.error('Failed to get version:', err);
		}
	});

	async function checkForUpdates() {
		isChecking = true;
		errorTitle = '';
		errorMessage = '';
		updateInfo = null;

		try {
			updateInfo = await CheckForUpdates();
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
</script>

<svelte:head>
	<title>{m['about.about']()}</title>
	<meta name="description" content={m['about.description']()} />
</svelte:head>

<div class="max-w-4xl mx-auto px-4 py-4 h-full">
	<div class="card bg-base-200 shadow-lg h-full overflow-auto">
		<div class="card-body p-6 md:p-10 flex flex-col md:flex-row items-center gap-6 mb-4">
			<div class="shrink-0">
				<img src={logo} alt="App icon" class="w-32 h-32 object-cover" />
			</div>

			<div class="flex-1 min-w-0">
				<h1 class="text-2xl md:text-3xl font-semibold leading-tight">TrID UI</h1>
				<div class="text-sm text-muted mt-1">
					{m['about.version']()} <span class="font-medium select-all">{version}</span>
				</div>

				<p class="mt-4 text-sm text-base-content/80">{m['about.description']()}</p>

				<div class="mt-6 flex flex-wrap gap-3 items-center">
					<button class="btn btn-primary btn-sm" on:click={checkForUpdates} disabled={isChecking} aria-label={m['about.check_for_updates']()}>
						{#if isChecking}
							<RefreshCw class="h-4 w-4 animate-spin" />
							{m['about.checking_for_updates']()}
						{:else}
							<RefreshCw class="h-4 w-4" />
							{m['about.check_for_updates']()}
						{/if}
					</button>

					<a
						class="inline-flex items-center gap-2 btn btn-sm"
						href="https://github.com/JMcrafter26/TridUI"
						target="_blank"
						rel="noopener noreferrer" aria-label={m['about.view_on_github']()}
					>
						<span class="w-4 h-4 opacity-80">
							<Github />
						</span>
						<span>{m['about.view_on_github']()}</span>
					</a>
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
					<div class="alert alert-soft alert-info mt-4">
						<Download class="h-6 w-6" />
						<div class="flex-1">
							<h3 class="font-bold">{m['about.update_available']()}</h3>
							<div class="text-sm mt-1">
								<p>
									{m['about.new_version_available'](updateInfo.latestVersion)}
									{#if updateInfo.publishedAt}
										({m['about.released']()} {formatDate(updateInfo.publishedAt)})
									{/if}
								</p>
							</div>
							{#if updateInfo.releaseNotes}
								<details class="mt-2">
									<summary class="cursor-pointer font-semibold text-sm" 
										>{m['about.release_notes']()}</summary
									>
									<div
										class="mt-2 p-3 bg-base-100 rounded-lg text-sm whitespace-pre-wrap max-h-48 overflow-y-auto"
									>
										{updateInfo.releaseNotes}
									</div>
								</details>
							{/if}
						</div>
						<button class="btn btn-sm btn-primary" on:click={openReleaseURL} aria-label={m['about.download']()}>
							<Download class="h-4 w-4" />
							{m['about.download']()}
						</button>
					</div>
				{:else}
					<div class="alert alert-soft alert-success mt-4 fixed bottom-3 md:bottom-6 max-w-md ">
						<CircleCheck class="h-5 w-5 my-auto" />
						<span>{m['about.up_to_date']()}</span>
						<button class="btn btn-ghost btn-sm rounded-full p-1 my-auto hover:bg-success/10" on:click={() => (updateInfo = null)}>
							<X />
						</button>
					</div>
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
					<p>{m['about.made_with_hearts_by']()}</p>
				</div>
			</div>
		</div>
	</div>
</div>
