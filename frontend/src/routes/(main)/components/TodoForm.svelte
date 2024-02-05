<script lang="ts">
	import Modal from '$lib/components/Modal.svelte';
	import Button from '$lib/components/button/Button.svelte';
	import Input from '$lib/components/input/Input.svelte';
	import Select from '$lib/components/input/Select.svelte';
	import { enhance } from '$app/forms';
	import { createEventDispatcher } from 'svelte';
	import type { Todo, TodoStatus } from '$lib/types/todo';

	export let open: boolean = false;
	export let defaultStatus: TodoStatus;
	export let selectedTodo: Todo | null;

	let formData = {
		title: '',
		description: '',
		status: defaultStatus
	};

	const dispatch = createEventDispatcher();

	$: {
		if (selectedTodo) {
			formData = {
				title: selectedTodo.title,
				description: selectedTodo.description,
				status: selectedTodo.status
			};
		} else {
			formData = {
				title: '',
				description: '',
				status: defaultStatus
			};
		}
	}

	const handleClose = () => {
		dispatch('toggle');
	};

	const statusOptions = [
		{
			label: 'Backlog',
			value: 'backlog'
		},
		{
			label: 'Pending',
			value: 'pending'
		},
		{
			label: 'In Progress',
			value: 'in-progress'
		},
		{
			label: 'Done',
			value: 'done'
		}
	];
</script>

<Modal {open} on:click={handleClose}>
	<div class="flex-1 p-6" on:click|stopPropagation role="none">
		<form
			class="flex flex-col gap-2"
			method="POST"
			action={selectedTodo ? '?/edit' : '?/create'}
			use:enhance={() =>
				async ({ result, update }) => {
					if (result.type === 'success') {
						dispatch('toggle');
					}
					setTimeout(() => {
						update();
					}, 100);
				}}
		>
			<input type="hidden" name="id" value={selectedTodo?.id} />
			<Input label="Title" name="title" value={formData.title} required />
			<Input label="Description" name="description" value={formData.description} />
			<Select name="status" label="Status" value={formData.status}>
				{#each statusOptions as option}
					<option value={option.value} selected={option.value === defaultStatus}
						>{option.label}</option
					>
				{/each}
			</Select>
			<div class="pt-6">
				<Button label="Submit" type="submit" />
			</div>
		</form>
	</div>
</Modal>
