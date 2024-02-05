import apiClient from '$lib/api/client';
import getTodoList from '$lib/api/todo/getTodoList';
import { TODO_STATUS } from '$lib/const';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ parent, fetch, depends }) => {
	if (typeof window !== 'undefined') {
		const parentData = await parent();
		apiClient.setToken(parentData.user.token);
	}
	apiClient.setFetch(fetch);
	const [backlog, pending, inProgress, done] = await Promise.all(
		TODO_STATUS.map((status) =>
			getTodoList({ status }).then((res) => {
				if (res.success) {
					return res.data;
				}
				return [];
			})
		)
	);
	depends('todo:list');

	return {
		todoList: {
			backlog,
			pending,
			'in-progress': inProgress,
			done
		}
	};
};
