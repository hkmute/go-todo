<script lang="ts">
	import TodoCard from '$lib/components/card/TodoCard.svelte';

	export let title: string;
	export let color: 'gray' | 'red' | 'yellow' | 'green';

	const todos = [
		{
			title: 'Todo 1',
			description: 'Description 1'
		},
		{
			title: 'Todo 2',
			description: 'Description 2'
		},
		{
			title: 'Todo 3',
			description: 'Description 3'
		},
		{
			title: 'Todo 4',
			description: 'Description 4'
		},
		{
			title: 'Todo 4',
			description: 'Description 4'
		},
		{
			title: 'Todo 4',
			description: 'Description 4'
		},
		{
			title: 'Todo 4',
			description: 'Description 4'
		},
		{
			title: 'Todo 4',
			description: 'Description 4'
		},
		{
			title: 'Todo 4',
			description: 'Description 4'
		},
		{
			title: 'Todo 4',
			description: 'Description 4'
		},
		{
			title: 'Todo 4',
			description: 'Description 4'
		}
	];

	const colors = {
		gray: 'border-gray-600/80',
		red: 'border-red-600/80',
		yellow: 'border-yellow-500/80',
		green: 'border-green-600/80'
	};
</script>

<div class="flex-1 rounded-2xl bg-indigo-100 p-2 pb-0 overflow-hidden flex flex-col">
	<div
		class={`mx-2 flex items-center justify-between border-b-4 pb-2.5 pt-2 text-lg font-semibold ${colors[color]}`}
	>
		<div class="px-2">{title}</div>
		<div class="px-2">+</div>
	</div>
	<div class="flex flex-col gap-4 px-2 py-4 overflow-auto">
		{#each todos as todo}
			<TodoCard title={todo.title} description={todo.description} />
		{/each}
	</div>
</div>
