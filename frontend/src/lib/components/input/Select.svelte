<script lang="ts">
	import ChervonDownIcon from '../icon/ChervonDownIcon.svelte';

	export let name: string;
	export let label: string;
	export let value: string = '';
</script>

<div class="flex flex-col gap-1">
	<label for={name}>{label}</label>
	<div class="relative">
		<select
			class="w-full appearance-none rounded-lg bg-slate-50 px-3 py-2"
			id={name}
			{name}
			bind:value
		>
			<slot />
		</select>
		<div class="icon pointer-events-none">
			<ChervonDownIcon size={16} />
		</div>
	</div>
</div>

<style>
	.icon {
		position: absolute;
		top: 50%;
		right: 0;
		transform: translateX(-50%) translateY(-50%);
	}
</style>
