import getMe from '$lib/api/user/getMe';
import type { LayoutServerLoad } from './$types';
import apiClient from '$lib/api/client';
import { redirect } from '@sveltejs/kit';

export const load: LayoutServerLoad = async ({ cookies, fetch }) => {
	apiClient.setFetch(fetch);
	const token = cookies.get('token');
	if (token) {
		const result = await getMe();
		if (result.success) {
			redirect(303, '/');
		}
        cookies.delete('token', {path: '/'});
	}
};
