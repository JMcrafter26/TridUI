<script lang="ts">
	import { onMount } from 'svelte';
	import { m } from '$lib/paraglide/messages.js';
	import { updateAvailable } from '$lib/stores/updateStore';

	import { House, Settings, Info, Menu, X, Minus, Pin } from '@lucide/svelte';
	import { WindowSetAlwaysOnTop, WindowMinimise, Quit } from '../../wailsjs/runtime/runtime.js';

	// Helper to get settings from localStorage
	function getSetting(key: string): unknown | null {
		const settingsData = localStorage.getItem('_trid_settings_');
		if (!settingsData) return null;

		const settings: Record<string, unknown> = JSON.parse(settingsData);
		return Object.prototype.hasOwnProperty.call(settings, key) ? settings[key] : null;
	}

	// pin button
	let isPinned = false;
	function togglePin() {
		isPinned = !isPinned;
		// Native function to set window always on top
		WindowSetAlwaysOnTop(isPinned);
	}

	onMount(() => {
		// Check if app should start pinned
		const startPinned = getSetting('startPinned');
		if (startPinned === true) {
			isPinned = true;
			WindowSetAlwaysOnTop(true);
		}
	});
</script>

<div
	class="fixed top-1 left-1/2 -translate-x-1/2 z-20 transition-all duration-300 ease-out mt-0.5"
	aria-hidden="true"
>
	<div
		class="group relative flex items-center justify-center h-6 w-20 rounded-full cursor-grab transition-all duration-300 ease-out hover:w-44 hover:h-8 active:w-44 active:h-8"
		style="--wails-draggable:drag"
	>
		<!-- background pill with faint red/yellow tint on left -->
		<div
			class="absolute inset-0 rounded-full bg-linear-to-r from-primary/20 via-secondary/15 via-40% to-accent/20 transition-all duration-300 ease-out group-hover:from-primary/70 group-hover:via-secondary/60 group-hover:to-accent/70 group-active:from-primary/70 group-active:via-secondary/60 group-active:to-accent/70"
		></div>

		<!-- the continuous bar that splits on hover -->
		<div
			class="absolute left-1/2 -translate-x-1/2 w-16 h-1.5 rounded-full bg-gray-300/50 pointer-events-none transition-all duration-200 ease-out group-hover:opacity-0"
		></div>

		<!-- left portion of split bar (hidden initially, appears on hover) -->
		<div
			class="absolute right-2 w-0 h-0 rounded-full bg-gray-300/80 pointer-events-none opacity-0 transition-all duration-300 ease-out group-hover:w-20 group-hover:h-4 group-hover:opacity-100 group-active:w-20 group-active:h-4 group-active:opacity-100"
		></div>

		<!-- traffic buttons that form from the right side of the bar -->
		<div
			class="absolute left-1/2 -translate-x-1/2 flex items-center gap-0 opacity-0 transition-all duration-300 ease-out group-hover:gap-1.5 group-hover:left-3 group-hover:translate-x-0 group-hover:opacity-100 group-active:gap-1.5 group-active:left-3 group-active:translate-x-0 group-active:opacity-100 pointer-events-auto"
		>
			<!-- close - morphs from bar segment -->
			<button
				on:click={Quit}
				class="flex items-center justify-center w-0 h-0 rounded-full bg-primary ring-1 ring-black/10 shadow-sm transition-all duration-300 ease-out group-hover:w-5 group-hover:h-5 group-active:w-5 group-active:h-5 overflow-hidden pointer-events-auto cursor-pointer hover:translate-y-0.5"
				title={m['header.close']()}
				aria-label={m['header.close']()}
			>
				<X
					class="w-0 h-0 text-primary-content opacity-0 transition-all duration-300 ease-out group-hover:w-3.5 group-hover:h-3.5 group-hover:opacity-80 group-active:w-3.5 group-active:h-3.5 group-active:opacity-80 pointer-events-none"
					strokeWidth={2.5}
				/>
			</button>
			<!-- minimize - morphs from bar segment -->
			<button
				on:click={WindowMinimise}
				class="flex items-center justify-center w-0 h-0 rounded-full bg-secondary ring-1 ring-black/10 shadow-sm transition-all duration-300 ease-out delay-50 group-hover:w-5 group-hover:h-5 group-hover:delay-0 group-active:w-5 group-active:h-5 group-active:delay-0 overflow-hidden pointer-events-auto cursor-pointer hover:translate-y-0.5"
				title={m['header.minimize']()}
				aria-label={m['header.minimize']()}
			>
				<Minus
					class="w-0 h-0 text-secondary-content opacity-0 transition-all duration-300 ease-out group-hover:w-3.5 group-hover:h-3.5 group-hover:opacity-80 group-active:w-3.5 group-active:h-3.5 group-active:opacity-80 pointer-events-none"
					strokeWidth={2.5}
				/>
			</button>
			<!-- pin - morphs from bar segment -->
			<button
				on:click={togglePin}
				class="flex items-center justify-center w-0 h-0 rounded-full bg-accent ring-1 ring-black/10 shadow-sm transition-all duration-300 ease-out delay-100 group-hover:w-5 group-hover:h-5 group-hover:delay-0 group-active:w-5 group-active:h-5 group-active:delay-0 overflow-hidden pointer-events-auto cursor-pointer hover:translate-y-0.5"
				title={m['header.pin']()}
				aria-label={m['header.pin']()}
			>
				<Pin
					class="w-0 h-0 text-accent-content opacity-0 transition-all duration-300 ease-out group-hover:w-3.5 group-hover:h-3.5 group-hover:opacity-80 group-active:w-3.5 group-active:h-3.5 group-active:opacity-80 pointer-events-none"
					strokeWidth={2.5}
					fill={isPinned ? 'currentColor' : 'none'}
				/>
			</button>
		</div>
	</div>
</div>

<div class="fab">
	<div
		tabindex="0"
		role="button"
		class="btn btn-soft btn-lg btn-circle btn-primary {$updateAvailable &&
		$updateAvailable.updateAvailable
			? 'indicator'
			: ''}"
	>
		{#if $updateAvailable && $updateAvailable.updateAvailable}
			<span class="indicator-item badge badge-secondary badge-xs mt-1.5 mr-1.5"></span>
		{/if}
		<Menu />
	</div>

	<!-- close button should not be focusable so it can close the FAB when clicked. It's just a visual placeholder -->
	<div class="fab-close">
		<span class="btn btn-circle btn-lg btn-primary"><X /></span>
	</div>

	<!-- buttons that show up when FAB is open -->
	<a
		class="btn btn-lg btn-circle {$updateAvailable && $updateAvailable.updateAvailable
			? 'indicator'
			: ''}"
		href="/about"
		title={m['header.about']()}
		aria-label={m['header.about']()}
	>
		{#if $updateAvailable && $updateAvailable.updateAvailable}
			<span class="indicator-item badge badge-secondary badge-xs mt-1.5 mr-1.5"></span>
		{/if}
		<Info />
	</a>
	<a
		class="btn btn-lg btn-circle"
		href="/settings"
		title={m['header.settings']()}
		aria-label={m['header.settings']()}
	>
		<Settings />
	</a>
	<a
		class="btn btn-lg btn-circle"
		href="/"
		title={m['header.home']()}
		aria-label={m['header.home']()}
	>
		<House />
	</a>
</div>

<style>
	@keyframes dropIn {
		0% {
			transform: translateY(-10px) scale(0);
			opacity: 0;
		}
		50% {
			transform: translateY(2px) scale(1.1);
		}
		100% {
			transform: translateY(0) scale(1);
			opacity: 1;
		}
	}
</style>
