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
)
