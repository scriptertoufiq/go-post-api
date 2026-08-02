package model

import "time"

type (
	UserModel struct {
		ID        int
		Name      string
		Username  string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	RefreshTokenModel struct {
		ID        int
		UserID    int64
		Token     string
		ExpiresAt time.Time
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
