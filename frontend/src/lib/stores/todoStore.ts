import type { TodoStatus } from '$lib/types/todo';
import { writable } from 'svelte/store';

type TodoStore = {
	defaultStatus: TodoStatus;
};

const todoStore = writable<TodoStore>({
	defaultStatus: 'backlog'
});

export default todoStore;
