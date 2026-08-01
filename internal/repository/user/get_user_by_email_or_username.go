package user

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
)

func (r *userRepository) GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error) {
	query := "SELECT id, name, email, username, password, created_at, updated_at FROM users WHERE email = ? OR username = ?"
	row := r.db.QueryRowContext(ctx, query, email, username)
	var result model.UserModel
	err := row.Scan(&result.ID, &result.Name, &result.Email, &result.Username, &result.Password, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err // Some other error occurred
	}
	return &result, nil
}
