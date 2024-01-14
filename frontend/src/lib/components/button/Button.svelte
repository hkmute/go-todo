<script lang="ts">
	import type { HTMLButtonAttributes } from 'svelte/elements';

	export let label: string;
	export let type: HTMLButtonAttributes['type'] = 'button';
	export let variant: 'contained' | 'outlined' | 'text' = 'contained';

	const classes = {
		contained: 'bg-indigo-200 hover:bg-indigo-300',
		outlined: 'border border-indigo-200 hover:border-indigo-300',
		text: 'hover:underline',
	};
</script>

<button
	class="w-full rounded-full px-4 py-2 font-semibold text-indigo-900 transition-colors {classes[variant]}"
	{type}
	on:click
>
	{label}
</button>
