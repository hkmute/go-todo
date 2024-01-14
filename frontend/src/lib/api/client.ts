import { PUBLIC_API_HOST } from '$env/static/public';

type ApiSuccessResponse<TData> = {
	success: true;
	data: TData;
};

type ApiErrorResponse = {
	success: false;
	message: string;
};

class ApiClient {
	private baseUrl: string = PUBLIC_API_HOST;

	constructor(private appFetch: typeof fetch = fetch) {}

	get = async <TData = undefined>(url: string, params?: Record<string, unknown>) => {
		try {
			const fetchUrl = new URL(url, this.baseUrl);
			if (params) {
				Object.keys(params).forEach((key) =>
					fetchUrl.searchParams.append(key, JSON.stringify(params[key]))
				);
			}
			const res = await this.appFetch(fetchUrl.toString());
			if (!res.ok) {
				return res.json() as Promise<ApiErrorResponse>;
			}

			return res.json() as Promise<ApiSuccessResponse<TData>>;
		} catch (error) {
			console.error(error);
			return { success: false, message: 'Something went wrong' } as const;
		}
	};

	post = async <TData = undefined>(url: string, body?: Record<string, unknown>) => {
		try {
			const fetchUrl = new URL(url, this.baseUrl);
			const res = await this.appFetch(fetchUrl, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify(body)
			});
			if (!res.ok) {
				return res.json() as Promise<ApiErrorResponse>;
			}

			return res.json() as Promise<ApiSuccessResponse<TData>>;
		} catch (error) {
			console.error(error);
			return { success: false, message: 'Something went wrong' } as const;
		}
	};

	put = async <TData = undefined>(url: string, body?: Record<string, unknown>) => {
		try {
			const fetchUrl = new URL(url, this.baseUrl);
			const res = await this.appFetch(fetchUrl, {
				method: 'PUT',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify(body)
			});

			if (!res.ok) {
				return res.json() as Promise<ApiErrorResponse>;
			}

			return res.json() as Promise<ApiSuccessResponse<TData>>;
		} catch (error) {
			console.error(error);
			return { success: false, message: 'Something went wrong' } as const;
		}
	};

	delete = async <TData = undefined>(url: string) => {
		try {
			const fetchUrl = new URL(url, this.baseUrl);
			const res = await this.appFetch(fetchUrl, {
				method: 'DELETE',
				headers: {
					'Content-Type': 'application/json',
				}
			});

			if (!res.ok) {
				return res.json() as Promise<ApiErrorResponse>;
			}

			return res.json() as Promise<ApiSuccessResponse<TData>>;
		} catch (error) {
			console.error(error);
			return { success: false, message: 'Something went wrong' } as const;
		}
	};

	setFetch(appFetch: typeof fetch) {
		this.appFetch = appFetch;
	}
}

const apiClient = new ApiClient();

export default apiClient;
