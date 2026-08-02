package dto

type (
	RegisterUserRequest struct {
		Name            string `json:"name" validate:"required"`
		Username        string `json:"username" validate:"required,min=3,max=20"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required,min=6"`
		PasswordConfirm string `json:"password_confirmation" validate:"required,eqfield=Password"`
	}

	RegisterUserResponse struct {
		ID string `json:"user_id"`
	}
)
