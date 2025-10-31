<script lang="ts">
	import Header from './Header.svelte';
	import { page } from '$app/stores';
	import '../app.css';
	let { children } = $props();
</script>

<div class="app h-screen flex flex-col">
	<Header />
 
	<main class="flex-1 overflow-auto">
		{@render children()}
	</main>
</div>
