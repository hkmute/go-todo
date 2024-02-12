import type { Todo, TodoStatus } from '$lib/types/todo';
import apiClient from '../client';

type UpdateTodoParams = {
	id: number;
	title: string;
	description?: string;
	status: TodoStatus;
};

const updateTodo = async ({ id, title, description, status }: UpdateTodoParams) => {
	return apiClient.put<Todo>(`/todo/${id}`, { title, description, status });
};

export default updateTodo;
