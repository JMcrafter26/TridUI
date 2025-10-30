<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import logo from '$lib/images/app-icon.png';
	import { Github } from '@lucide/svelte';

	const version = '1.0.0';
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
				<h1 class="text-2xl md:text-3xl font-semibold leading-tight">TrID UI</h1>
				<div class="text-sm text-muted mt-1">{m["about.version"]()} <span class="font-medium">{version}</span></div>

				<p class="mt-4 text-sm text-base-content/80 select-text">{m["about.description"]()}</p>

				<div class="mt-6 flex flex-wrap gap-3 items-center">

					<button class="btn btn-primary btn-sm" title="Check for updates (not implemented)">{m["about.check_for_updates"]()}</button>

					<a class="inline-flex items-center gap-2 btn btn-ghost btn-sm" href="https://github.com/" target="_blank" rel="noopener noreferrer">
						<span class="w-4 h-4 opacity-80">
							<Github />
						</span>
						<span>{m["about.view_on_github"]()}</span>
					</a>

				</div>
			</div>

			<div class="divider md:divider-horizontal"></div>

			<div class="w-full md:w-56">
				<h3 class="text-sm font-medium">{m["about.about_and_credits"]()}</h3>
				<p class="text-xs text-base-content/70 mt-2 select-text">{@html m["about.made_by"]()}</p>

				<div class="mt-4 text-xs text-base-content/70">
					<div>{m["about.license"]()}: GNU AGPL v3.0</div>
					<div class="mt-2">{m["about.built_with"]()}</div>
				</div>
			</div>
		</div>
	</div>
</div>
