package todo

type GetTodoListParams struct {
	Offset int `form:"offset" binding:"number"`
	Limit  int `form:"limit" binding:"number"`
}

type NewTodo struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description"`
	Status      string `form:"status" binding:"required,oneof=backlog pending in-progress done"`
	UserId      int
}

type EditTodo struct {
	Title       *string `form:"title"`
	Description *string `form:"description"`
	Status      string  `form:"status" binding:"omitempty,oneof=backlog pending in-progress done"`
	UserId      int
}
