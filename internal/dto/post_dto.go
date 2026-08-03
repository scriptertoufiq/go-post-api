package dto

type (
	CreatePostRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	CreatePostResponse struct {
		ID int64 `json:"id"`
	}
)
