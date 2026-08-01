package dto

type (
	RegisterUserRequest struct {
		Name     string `json:"name" binding:"required"`
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	RegisterUserResponse struct {
		ID string `json:"user_id"`
	}
)
