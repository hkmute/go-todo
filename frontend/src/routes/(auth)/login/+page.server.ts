import type { Actions } from './$types';
import apiClient from '$lib/api/client';
import { fail, redirect } from '@sveltejs/kit';

export const actions = {
	default: async ({ request, cookies }) => {
		const data = await request.formData();
		const username = data.get('username');
		const password = data.get('password');

		const result = await apiClient.post<string>('/user/login', { username, password });
		if (result.success) {
			cookies.set('token', result.data, { path: '/' });
			redirect(303, '/')
		}
		return fail(401, result);
	}
} satisfies Actions;
