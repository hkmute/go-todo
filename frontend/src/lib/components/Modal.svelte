<script lang="ts">
	import { fade } from 'svelte/transition';

	export let open: boolean = false;
</script>

{#if open}
	<div
		transition:fade={{ duration: 100 }}
		class="fixed left-0 top-0 flex h-dvh w-dvw items-center justify-center bg-gray-900/10 transition-opacity"
		on:click
		role="presentation"
	>
		<div class="flex min-w-[400px] flex-col rounded-xl bg-indigo-100">
			<slot />
		</div>
	</div>
{/if}
