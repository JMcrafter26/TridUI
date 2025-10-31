<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import { WindowSetSize } from '../../../wailsjs/runtime/runtime';
</script>

<svelte:head>
	<title>{m['settings.settings']()}</title>
	<meta name="description" content={m['settings.description']()} />
</svelte:head>

<div class="max-w-4xl mx-auto px-4 py-4 h-full">
	<div class="card bg-base-200 shadow-lg h-full">
		<div class="card-body p-6 md:p-10 flex flex-col md:flex-row gap-6">
			<h2 class="card-title text-xl">{m['settings.settings']()}</h2>
			<p class="text-sm opacity-70 hidden">{m['settings.description']()}</p>

			<div class="mt-4">
				<!-- Settings content goes here -->
				<p class="text-center text-sm opacity-50">Settings page is under construction.</p>
			</div>

			<!-- debug -->
			<details class="collapse border-base-300 border">
				<summary class="collapse-title text-sm opacity-70">Debug Tools</summary>
				<div class="collapse-content text-sm">
					<button class="btn btn-primary btn-dash" on:click={() => WindowSetSize(500, 400)}
					>Reset window size</button>
				</div>
			</details>
		
		</div>
	</div>
</div>
