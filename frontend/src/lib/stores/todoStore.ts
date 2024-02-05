import type { Todo, TodoStatus } from '$lib/types/todo';
import { writable } from 'svelte/store';

type TodoStates = {
	todoLists: Record<TodoStatus, Todo[]>;
	defaultStatus: TodoStatus;
};

const todoStates = writable<TodoStates>({
	todoLists: {
		backlog: [],
		pending: [],
		'in-progress': [],
		done: []
	},
	defaultStatus: 'backlog'
});

const moveTodo = (todo: Todo, newStatus: TodoStatus, dropPosition: number) => {
	todoStates.update((state) => {
		if (todo.status === newStatus) {
			const originalIndex = state.todoLists[newStatus].findIndex((item) => item.id === todo.id);
			state.todoLists[newStatus].splice(originalIndex, 1);
			state.todoLists[newStatus].splice(
				originalIndex < dropPosition ? dropPosition - 1 : dropPosition,
				0,
				todo
			);

			return {
				...state,
				todoLists: {
					...state.todoLists,
					[newStatus]: state.todoLists[newStatus]
				}
			};
		}
		const oldStatus = todo.status;
		todo.status = newStatus;
		const oldTodoList = state.todoLists[oldStatus].filter((item) => item.id !== todo.id);
		const newTodoList = state.todoLists[newStatus];
		newTodoList.splice(dropPosition, 0, todo);
		return {
			...state,
			todoLists: {
				...state.todoLists,
				[oldStatus]: oldTodoList,
				[newStatus]: newTodoList
			}
		};
	});
};

const todoStore = {
	...todoStates,
	moveTodo
};

export default todoStore;
