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
 
	<footer>
		<p>&copy; 2025 TridscanUI</p>
		<p>{m.made_by_footer()}</p>
	</footer>
</div>
