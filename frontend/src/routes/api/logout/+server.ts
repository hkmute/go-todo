import { redirect, type RequestEvent } from '@sveltejs/kit';

export const POST = async ({ cookies }: RequestEvent) => {
	cookies.delete('token', { path: '/' });
	return redirect(303, '/login');
};
