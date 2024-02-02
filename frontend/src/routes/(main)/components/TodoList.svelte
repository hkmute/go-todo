<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import TodoCard from '$lib/components/card/TodoCard.svelte';
	import PlusIcon from '$lib/components/icon/PlusIcon.svelte';
	import type { Todo, TodoStatus } from '$lib/types/todo';
	import todoStore from '$lib/stores/todoStore';

	export let title: string;
	export let color: 'gray' | 'red' | 'yellow' | 'green';
	export let items: Todo[];
	export let status: TodoStatus;

	const dispatch = createEventDispatcher();

	const handleOpen = () => {
		todoStore.update((store) => {
			return {
				...store,
				defaultStatus: status,
			}
		});
		dispatch('toggle');
	};

	const colors = {
		gray: 'border-gray-600/80',
		red: 'border-red-600/80',
		yellow: 'border-yellow-500/80',
		green: 'border-green-600/80'
	};
</script>

<div class="flex flex-1 flex-col overflow-hidden rounded-2xl bg-indigo-100 pb-0">
	<div
		class={`mx-4 flex items-center justify-between border-b-4 pb-2.5 pt-4 text-lg font-semibold ${colors[color]}`}
	>
		<div class="px-2">{title}</div>
		<button
			class="flex h-8 w-8 min-w-8 items-center justify-center rounded-full transition-colors duration-100 hover:bg-indigo-200"
			on:click={handleOpen}
		>
			<PlusIcon size={16} />
		</button>
	</div>
	<div class="flex flex-col gap-4 overflow-auto px-4 py-4">
		{#each items as todo}
			<TodoCard todo={todo} on:toggle />
		{/each}
	</div>
</div>
