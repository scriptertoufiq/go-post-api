package post

import (
	"context"
	"go-tweets/internal/model"
)

func (r *postRepository) StorePost(ctx context.Context, post *model.PostModel) (int64, error) {
	query := `INSERT INTO posts (user_id, title, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.ExecContext(ctx, query, post.UserID, post.Title, post.Content, post.CreatedAt, post.UpdatedAt)
	if err != nil {
		return 0, err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return postID, nil
}
