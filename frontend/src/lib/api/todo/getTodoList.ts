import type { Todo, TodoStatus } from '$lib/types/todo';
import apiClient from '../client';

type GetTodoListParams = {
	limit?: number;
	offset?: number;
	status?: TodoStatus;
};

const getTodoList = async (params?: GetTodoListParams) => {
	return apiClient.get<Todo[]>('/todo/', params);
};

export default getTodoList;
