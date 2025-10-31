<script lang="ts">
	import { onMount } from 'svelte';
	import { m } from '$lib/paraglide/messages.js';
	import { WindowSetSize, EventsOn, EventsOff } from '../../../wailsjs/runtime/runtime';
	import {
		CheckDefinitionsExist,
		CheckForDefsUpdates,
		UpdateDefinitions,
		GetDefinitionsPath,
		OpenAppDir
	} from '../../../wailsjs/go/main/App';
	import { Download, RefreshCw, FolderOpen, CheckCircle, AlertCircle, Info } from '@lucide/svelte';

	let definitionsExist = false;
	let definitionsPath = '';
	let updateInfo: any | null = null;
	let isCheckingUpdates = false;
	let isUpdating = false;
	let updateProgress = 0;
	let updateMessage = '';
	let updateError = '';

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
				<div class="divider">TrID Definitions</div>

				<!-- Definitions Status -->
				<div class="alert {definitionsExist ? 'alert-success' : 'alert-warning'}">
					{#if definitionsExist}
						<CheckCircle class="h-5 w-5" />
						<div class="flex-1">
							<h3 class="font-bold">Definitions Installed</h3>
							<div class="text-sm">
								{#if updateInfo}
									<div class="mt-1 space-y-1">
										<div>Definitions: {updateInfo.defsCount.toLocaleString()} file types</div>
										{#if updateInfo.lastUpdated}
											<div>Last updated: {updateInfo.lastUpdated}</div>
										{/if}
										{#if updateInfo.currentMD5 && updateInfo.currentMD5 !== 'none'}
											<div class="text-xs opacity-70">MD5: {updateInfo.currentMD5}</div>
										{/if}
									</div>
								{:else}
									Located at: {definitionsPath}
								{/if}
							</div>
						</div>
					{:else}
						<AlertCircle class="h-5 w-5" />
						<div class="flex-1">
							<h3 class="font-bold">Definitions Not Found</h3>
							<div class="text-sm">
								Download the definitions file to enable file identification.
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
										{updateInfo.isUpToDate ? 'Up to date' : 'Update available'}
									</span>
								</div>
								{#if !updateInfo.isUpToDate}
									<span class="badge badge-primary">New version</span>
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
						<AlertCircle class="h-5 w-5" />
						<div>
							<h3 class="font-bold">Error</h3>
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
					>
						{#if isCheckingUpdates}
							<RefreshCw class="h-4 w-4 animate-spin" />
							Checking...
						{:else}
							<RefreshCw class="h-4 w-4" />
							Check for Updates
						{/if}
					</button>

					{#if !definitionsExist || (updateInfo && !updateInfo.isUpToDate)}
						<button class="btn btn-success" on:click={downloadUpdates} disabled={isUpdating}>
							{#if isUpdating}
								<Download class="h-4 w-4 animate-bounce" />
								Updating...
							{:else}
								<Download class="h-4 w-4" />
								{definitionsExist ? 'Update Definitions' : 'Download Definitions'}
							{/if}
						</button>
					{/if}

					<button class="btn btn-ghost" on:click={openAppDirectory}>
						<FolderOpen class="h-4 w-4" />
						Open App Directory
					</button>
				</div>

				<!-- Info Box -->
				<div class="text-xs opacity-70 mt-4 p-3 bg-base-300 rounded">
					<p class="mb-1">
						<strong>Note:</strong> TrID definitions are updated regularly by Marco Pontello to support
						new file types.
					</p>
					<p>
						Definitions are downloaded from:
						<a
							href="http://mark0.net/soft-trid-e.html"
							target="_blank"
							rel="noopener noreferrer"
							class="link">mark0.net</a
						>
					</p>
				</div>
			</div>

			<!-- Debug Section -->
			<details class="collapse border-base-300 border mt-6">
				<summary class="collapse-title text-sm opacity-70">Debug Tools</summary>
				<div class="collapse-content text-sm space-y-2">
					<button class="btn btn-sm btn-primary" on:click={() => WindowSetSize(500, 400)}>
						Reset window size
					</button>
					{#if definitionsPath}
						<div class="text-xs opacity-70 break-all">
							<strong>Definitions path:</strong>
							{definitionsPath}
						</div>
					{/if}
				</div>
			</details>
		</div>
	</div>
</div>
