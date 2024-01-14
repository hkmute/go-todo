import apiClient from '../client';

type User = {
	id: string;
	username: string;
};

const getMe = async () => {
	return apiClient.get<User>('/user/me');
};

export default getMe;
