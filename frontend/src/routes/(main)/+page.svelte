<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/button/Button.svelte';
	import todoStore from '$lib/stores/todoStore';
	import type { Todo } from '$lib/types/todo';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import TodoForm from './components/TodoForm.svelte';
	import TodoList from './components/TodoList.svelte';

	export let data: PageData;

	$: todoStore.update((store) => {
		return { ...store, todoLists: data.todoList };
	});

	let open = false;
	let selectedTodo: Todo | null;

	const handleLogout = async () => {
		const result = await fetch('/api/logout', {
			method: 'POST'
		});
		if (result.redirected) {
			goto(result.url);
		}
	};

	const handleToggle = (v: CustomEvent<{ todo: Todo } | undefined>) => {
		open = !open;
		selectedTodo = v.detail?.todo ?? null;
	};
</script>

<div class="flex h-dvh max-h-dvh flex-col">
	<header class="flex items-center justify-between bg-indigo-100 px-4 py-2">
		<h1>Welcome back, {data.user.username}!</h1>
		<div>
			<Button label="Logout" on:click={handleLogout} />
		</div>
	</header>

	<main class="flex flex-1 gap-4 overflow-hidden bg-indigo-50 p-4">
		<TodoList
			on:toggle={handleToggle}
			title="Backlog"
			color="gray"
			items={$todoStore.todoLists.backlog}
			status="backlog"
		/>
		<TodoList
			on:toggle={handleToggle}
			title="Pending"
			color="red"
			items={$todoStore.todoLists.pending}
			status="pending"
		/>
		<TodoList
			on:toggle={handleToggle}
			title="In Progress"
			color="yellow"
			items={$todoStore.todoLists['in-progress']}
			status="in-progress"
		/>
		<TodoList
			on:toggle={handleToggle}
			title="Done"
			color="green"
			items={$todoStore.todoLists.done}
			status="done"
		/>
		<TodoForm
			{open}
			on:toggle={handleToggle}
			defaultStatus={$todoStore.defaultStatus}
			{selectedTodo}
		/>
	</main>
</div>
