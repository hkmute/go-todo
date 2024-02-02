<script lang="ts">
	import { invalidate } from '$app/navigation';
	import deleteTodo from '$lib/api/todo/deleteTodo';
	import DeleteIcon from '$lib/components/icon/DeleteIcon.svelte';
	import type { Todo } from '$lib/types/todo';
	import { createEventDispatcher } from 'svelte';

	export let todo: Todo;

	const dispatch = createEventDispatcher();

	const handleClick = () => {
		dispatch('toggle', { todo });
	};

	const handleDelete = async () => {
		await deleteTodo({ id: todo.id });
		invalidate('todo:list');
	};
</script>

<button
	class="block flex-1 rounded-2xl bg-indigo-50/65 p-4 text-left transition-colors duration-100 hover:bg-indigo-50/100"
	on:click={handleClick}
>
	<div class="flex items-center justify-between">
		<div>{todo.title}</div>
		<button
			class="flex h-6 w-6 items-center justify-center rounded-full hover:bg-indigo-100"
			on:click|stopPropagation={handleDelete}
		>
			<DeleteIcon size={16} />
		</button>
	</div>
	{#if todo.description}
		<div class="py-1 text-sm">{todo.description}</div>
	{/if}
</button>
