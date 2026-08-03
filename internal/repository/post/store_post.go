package post

import (
	"context"
	"go-tweets/internal/model"
)

func (r *postRepository) StorePost(ctx context.Context, post *model.PostModel) (int64, error) {
	query := `INSERT INTO posts (user_id, content, created_at) VALUES ($1, $2, $3) RETURNING id`
	var postID int64
	err := r.db.QueryRowContext(ctx, query, post.UserID, post.Content, post.CreatedAt).Scan(&postID)
	if err != nil {
		return 0, err
	}
	return postID, nil
}
