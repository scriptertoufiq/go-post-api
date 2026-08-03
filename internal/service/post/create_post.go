package post

import (
	"context"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"net/http"
	"time"
)

func (s *postService) CreatePost(ctx context.Context, req *dto.CreatePostRequest) (int64, int, error) {
	now := time.Now()
	insertedId, err := s.postRepo.StorePost(ctx, &model.PostModel{
		UserID:    req.UserID,
		Content:   req.Content,
		CreatedAt: now,
	})
	if err != nil {
		return 0, 0, err
	}
	return insertedId, http.StatusCreated, nil
}
