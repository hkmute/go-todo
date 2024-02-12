import type { Todo, TodoStatus } from '$lib/types/todo';
import apiClient from '../client';

type UpdateTodoParams = {
	id: number;
	title: string;
	description?: string;
	status: TodoStatus;
	itemOrder: number;
};

const updateTodo = async ({ id, title, description, status, itemOrder }: UpdateTodoParams) => {
	return apiClient.put<Todo>(`/todo/${id}`, { title, description, status, itemOrder });
};

export default updateTodo;
