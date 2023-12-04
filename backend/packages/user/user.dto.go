package user

type GetUserParams struct {
	Id       *int    `uri:"id" binding:"omitempty,number"`
	Username *string `uri:"username" binding:"omitempty"`
}

type NewUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EditUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
