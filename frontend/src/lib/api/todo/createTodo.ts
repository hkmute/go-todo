import type { Todo, TodoStatus } from '$lib/types/todo';
import apiClient from '../client';

type CreateTodoParams = {
	title: string;
	description?: string;
	status: TodoStatus;
};

const createTodo = async ({ title, description, status }: CreateTodoParams) => {
	return apiClient.post<Todo>('/todo', { title, description, status });
};

export default createTodo;
