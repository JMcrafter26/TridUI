<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import logo from '$lib/images/app-icon.png';
	import { Github } from '@lucide/svelte';

	const appName = 'Tridscan UI';
	const version = '1.0.0';
	const description = `A lightweight desktop UI for Tridscan. Select or drop a file on the Home screen to start a local scan.`;
</script>

<svelte:head>
	<title>{m["about.about"]()}</title>
	<meta name="description" content={m["about.description"]()} />
</svelte:head>

<div class="max-w-4xl mx-auto px-4 py-12">
	<div class="card bg-base-200 shadow-lg">
		<div class="card-body p-6 md:p-10 flex flex-col md:flex-row items-center gap-6">
			<div class="shrink-0">
				<img src={logo} alt="App icon" class="w-32 h-32 object-cover" />
			</div>

			<div class="flex-1 min-w-0">
				<h1 class="text-2xl md:text-3xl font-semibold leading-tight">{appName}</h1>
				<div class="text-sm text-muted mt-1">Version <span class="font-medium">{version}</span></div>

				<p class="mt-4 text-sm text-base-content/80">{description}</p>

				<div class="mt-6 flex flex-wrap gap-3 items-center">
					<a class="btn btn-primary btn-sm" href="/" title="Go to Home">Open Home</a>

					<a class="btn btn-outline btn-sm" href="#" title="Check for updates (not implemented)">Check for updates</a>

					<a class="inline-flex items-center gap-2 btn btn-ghost btn-sm" href="https://github.com/" target="_blank" rel="noopener noreferrer">
						<span class="w-4 h-4 opacity-80">
							<Github />
						</span>
						<span>View on GitHub</span>
					</a>
				</div>
			</div>

			<div class="divider md:divider-horizontal"></div>

			<div class="w-full md:w-56">
				<h3 class="text-sm font-medium">About & Credits</h3>
				<p class="text-xs text-base-content/70 mt-2 select-text">{@html m["about.made_by"]()}</p>

				<div class="mt-4 text-xs text-base-content/70">
					<div>License: MIT</div>
					<div class="mt-2">Built with Wails + Svelte + Tailwind</div>
				</div>
			</div>
		</div>
	</div>
</div>
