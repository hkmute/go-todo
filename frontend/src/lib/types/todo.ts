export type Todo = {
	id: number;
	title: string;
	description: string;
	status: TodoStatus;
	userId: number;
	createdAt: string;
	updatedAt: string;
};

export type TodoStatus = 'backlog' | 'pending' | 'in-progress' | 'done';
