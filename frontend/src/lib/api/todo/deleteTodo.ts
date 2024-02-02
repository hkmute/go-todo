import apiClient from "../client";

type DeleteTodoParams = {
  id: number;
}

const deleteTodo = async ({ id }: DeleteTodoParams) => {
  return apiClient.delete<void>(`/todo/${id}`);
};

export default deleteTodo;