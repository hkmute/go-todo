import type { Todo, TodoStatus } from '$lib/types/todo';
import apiClient from '../client';

type CreateTodoParams = {
	id: number;
	title: string;
	description?: string;
	status: TodoStatus;
};

const createTodo = async ({ id, title, description, status }: CreateTodoParams) => {
	return apiClient.put<Todo>(`/todo/${id}`, { title, description, status });
};

export default createTodo;
