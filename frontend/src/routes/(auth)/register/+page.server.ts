import apiClient from '$lib/api/client';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions = {
	default: async ({request}) => {
		const data = await request.formData();
		const username = data.get('username');
		const password = data.get('password');
		const confirmPassword = data.get('confirm-password');

		if (password !== confirmPassword) {
			return fail(400, { message: 'Passwords do not match' });
		}

		const result = await apiClient.post<{ success: boolean; token?: string }>(
			'/user/register',
			{ username, password }
		);

		if (result.success) {
			redirect(303, '/login');
		}
		return fail(401, result);
		},
} satisfies Actions;