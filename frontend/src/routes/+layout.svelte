<script lang="ts">
	import Header from './Header.svelte';
	import { page } from '$app/stores';
	import '../app.css';
	let { children } = $props();
</script>

<div class="app">
	<Header />
 
	<main>
		{@render children()}
	</main>

	<footer class="footer sm:footer-horizontal footer-center fixed bottom-0 left-0 w-full py-3 shadow z-10">
		<div class="max-w-4xl mx-auto">
			{#if $page.url.pathname == '/about'}
				<p class="badge badge-ghost"><a href="/">&copy; 2025 TridscanUI</a><a href="https://go.jm26.net/github" target="_blank" rel="noopener noreferrer">by Cufiy</a></p>
			{:else}
				<p class="badge badge-ghost"><a href="/about">&copy; 2025 TridscanUI</a><a href="https://go.jm26.net/github" target="_blank" rel="noopener noreferrer">by Cufiy</a></p>
			{/if}
		</div>
	</footer>
</div>
