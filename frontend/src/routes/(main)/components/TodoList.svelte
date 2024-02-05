<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import TodoCard from '$lib/components/card/TodoCard.svelte';
	import PlusIcon from '$lib/components/icon/PlusIcon.svelte';
	import type { Todo, TodoStatus } from '$lib/types/todo';
	import todoStore from '$lib/stores/todoStore';
	import { crossfade } from 'svelte/transition';
	import { flip } from 'svelte/animate';

	export let title: string;
	export let color: 'gray' | 'red' | 'yellow' | 'green';
	export let items: Todo[];
	export let status: TodoStatus;

	let scrolledY = 0;
	let dropPosition = items.length;

	const dispatch = createEventDispatcher();
	const [send, receive] = crossfade({});

	const handleOpen = () => {
		todoStore.update((store) => {
			return {
				...store,
				defaultStatus: status
			};
		});
		dispatch('toggle');
	};

	const handleDrop = (e: DragEvent) => {
		if (!e.dataTransfer) return;
		const todo = JSON.parse(e.dataTransfer.getData('application/json'));
		todoStore.moveTodo(todo, status, dropPosition);
	};

	const handleDragover = (e: DragEvent) => {
		const currentTarget = e.currentTarget as HTMLDivElement;
		scrolledY = currentTarget.scrollTop;
	};

	const handleDragoverItem = (index: number) => (e: DragEvent) => {
		const target = e.target as HTMLDivElement;
		const middlePoint = target.offsetTop + target.offsetHeight / 2;
		if (e.clientY + scrolledY > middlePoint) {
			dropPosition = index + 1;
		} else {
			dropPosition = index;
		}
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
	<div
		class="flex h-full flex-col overflow-auto px-4 py-2"
		role="list"
		on:drop={handleDrop}
		on:dragover|preventDefault={handleDragover}
	>
		{#each items as todo, index (todo.id)}
			<div
				class="py-2"
				role="listitem"
				on:dragover|preventDefault={handleDragoverItem(index)}
				animate:flip={{ duration: 300 }}
				in:receive={{ key: todo.id }}
				out:send={{ key: todo.id }}
			>
				<TodoCard {todo} on:toggle />
			</div>
		{/each}
	</div>
</div>
