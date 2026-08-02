package user

import (
	"context"
	"go-tweets/internal/model"
)

func (r *userRepository) StoreRefreshToken(ctx context.Context, refreshToken *model.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (user_id, token, expires_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, refreshToken.UserID, refreshToken.Token, refreshToken.ExpiresAt, refreshToken.CreatedAt, refreshToken.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
