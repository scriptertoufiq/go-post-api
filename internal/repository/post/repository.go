package post

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
)

type PostRepository interface {
	StorePost(ctx context.Context, post *model.PostModel) (int64, error)
}

type postRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}
