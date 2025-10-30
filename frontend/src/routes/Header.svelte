<script lang="ts">
	import { page } from '$app/stores';
	import { House, Settings, Info, Menu, X } from '@lucide/svelte';
</script>

<!-- <header class="relative">
	<ul class="menu menu-horizontal bg-base-200 rounded-box fixed top-0 right-0 m-4 z-10 shadow-lg">
		<li>
			<a href="/" class:active={$page.url.pathname === '/'}>
				<House />
			</a>
		</li>
		<li>
			<a href="/settings" class:active={$page.url.pathname === '/settings'}>
				<Settings />
			</a>
		</li>
		<li>
			<a href="/about" class:active={$page.url.pathname === '/about'}>
				<Info />
			</a>
		</li>
	</ul>
</header> -->

<div class="fab">
  <div tabindex="0" role="button" class="btn btn-soft btn-lg btn-circle btn-primary"><Menu /></div>

  <!-- close button should not be focusable so it can close the FAB when clicked. It's just a visual placeholder -->
  <div class="fab-close">
    <span class="btn btn-circle btn-lg btn-primary"><X /></span>
  </div>

	<!-- buttons that show up when FAB is open -->
	<a class="btn btn-lg btn-circle" href="/about">
		<Info />
	</a>
	<a class="btn btn-lg btn-circle" href="/settings">
		<Settings />
	</a>
	<a class="btn btn-lg btn-circle" href="/">
		<House />
	</a>
</div>
