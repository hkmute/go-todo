import { fail } from '@sveltejs/kit';
import type { Actions } from './$types';
import createTodo from '$lib/api/todo/createTodo';
import type { TodoStatus } from '$lib/types/todo';
import updateTodo from '$lib/api/todo/updateTodo';

export const actions = {
	create: async ({ request }) => {
		const data = await request.formData();
		const title = data.get('title') as string;
		const description = data.get('description') as string | undefined;
		const status = data.get('status') as TodoStatus;
		const result = await createTodo({ title, description, status });
		if (result.success) {
			return result;
		}
		return fail(400, result);
	},
	edit: async ({ request }) => {
		const data = await request.formData();
		const id = parseInt(data.get('id') as string);
		const title = data.get('title') as string;
		const description = data.get('description') as string | undefined;
		const status = data.get('status') as TodoStatus;
		const itemOrder = parseInt(data.get('itemOrder') as string);
		const result = await updateTodo({ id, title, description, status, itemOrder });
		if (result.success) {
			return result;
		}
	},
	delete: async () => {}
} satisfies Actions;
