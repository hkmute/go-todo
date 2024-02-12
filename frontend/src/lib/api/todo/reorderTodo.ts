import type { Todo, TodoStatus } from '$lib/types/todo';
import apiClient from '../client';

type ReorderTodoParams = {
	id: number;
	status: TodoStatus;
	itemOrder: number;
};

const reorderTodo = async ({ id, status, itemOrder }: ReorderTodoParams) => {
	return apiClient.put<Todo>(`/todo/${id}/reorder`, { status, itemOrder });
};

export default reorderTodo;
