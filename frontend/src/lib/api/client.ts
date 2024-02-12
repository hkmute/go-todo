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
	private token: string | null = null;

	constructor(private appFetch: typeof fetch = fetch) {}

	get = async <TData = undefined>(url: string, params?: Record<string, unknown>) => {
		try {
			const fetchUrl = url.startsWith('http') ? new URL(url) : new URL(url, this.baseUrl);
			if (params) {
				Object.keys(params).forEach((key) =>
					fetchUrl.searchParams.append(key, params[key] as string)
				);
			}

			const res = await this.appFetch(fetchUrl, {
				headers: {
					Authorization: this.token ? `Bearer ${this.token}` : ''
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

	post = async <TData = undefined>(url: string, body?: Record<string, unknown>) => {
		try {
			const fetchUrl = url.startsWith('http') ? new URL(url) : new URL(url, this.baseUrl);
			const res = await this.appFetch(fetchUrl, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Authorization: this.token ? `Bearer ${this.token}` : ''
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
			const fetchUrl = url.startsWith('http') ? new URL(url) : new URL(url, this.baseUrl);
			const res = await this.appFetch(fetchUrl, {
				method: 'PUT',
				headers: {
					'Content-Type': 'application/json',
					Authorization: this.token ? `Bearer ${this.token}` : ''
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
			const fetchUrl = url.startsWith('http') ? new URL(url) : new URL(url, this.baseUrl);
			const res = await this.appFetch(fetchUrl, {
				method: 'DELETE',
				headers: {
					'Content-Type': 'application/json',
					Authorization: this.token ? `Bearer ${this.token}` : ''
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

	setToken(token: string | null) {
		if (typeof window === 'undefined') {
			throw new Error('Cannot set token on server side');
		}
		this.token = token;
	}
}

const apiClient = new ApiClient();

export default apiClient;
