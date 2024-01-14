import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';
import apiClient from '$lib/api/client';
import getMe from '$lib/api/user/getMe';

export const load: LayoutServerLoad = async ({ fetch, cookies }) => {
	apiClient.setFetch(fetch);
	const token = cookies.get('token');
	if (token) {
		const result = await getMe();

		if (result.success) {
			return result.data;
		}
	}
	cookies.delete('token', { path: '/' });
	redirect(303, '/login');
};
