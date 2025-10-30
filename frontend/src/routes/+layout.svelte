<script lang="ts">
	import Header from './Header.svelte';
	import '../app.css';
	import { setLocale } from '$lib/paraglide/runtime';
	import { m } from '$lib/paraglide/messages.js';


	let { children } = $props();
</script>

<div class="app">
	<Header />
 
	<main>
		{@render children()}
	</main>
 
	<footer class="footer sm:footer-horizontal footer-center fixed bottom-0 left-0 w-full bg-base-200/80 backdrop-blur py-3 shadow z-10">
		<div class="max-w-4xl mx-auto">
			<p>&copy; 2025 TridscanUI</p>
			<p>{@html m.made_by_footer()}</p>
		</div>
	</footer>
</div>
