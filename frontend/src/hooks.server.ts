import type { HandleFetch } from '@sveltejs/kit';

export const handleFetch: HandleFetch = async ({ event, request, fetch }) => {
	const token = event.cookies.get('token');
	if (token) {
		request.headers.set('Authorization', `Bearer ${token}`);
	}

	return fetch(request);
};
