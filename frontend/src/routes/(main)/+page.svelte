<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/button/Button.svelte';
	import TodoCard from '$lib/components/card/TodoCard.svelte';
	import type { PageData } from './$types';
	import TodoList from './components/TodoList.svelte';

	export let data: PageData;

	const handleLogout = async () => {
		const result = await fetch('/api/logout', {
			method: 'POST'
		});
		if (result.redirected) {
			goto(result.url);
		}
	};
</script>

<div class="flex max-h-dvh h-dvh flex-col">
	<header class="flex items-center justify-between bg-indigo-100 px-4 py-2">
		<h1>Welcome back, {data.username}!</h1>
		<div>
			<Button label="Logout" on:click={handleLogout} />
		</div>
	</header>

	<main class="flex flex-1 gap-4 bg-indigo-50 p-4 overflow-hidden">
		<TodoList title="Backlog" color="gray" />
		<TodoList title="Pending" color="red" />
		<TodoList title="In Progress" color="yellow" />
		<TodoList title="Done" color="green" />
	</main>
</div>
