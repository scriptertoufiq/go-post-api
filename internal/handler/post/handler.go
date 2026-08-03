package post

import (
	"go-tweets/internal/service/post"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api         *gin.Engine
	validate    *validator.Validate
	postService post.PostService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, postService post.PostService) *Handler {
	return &Handler{
		api:         api,
		validate:    validate,
		postService: postService,
	}
}

func (h *Handler) RouteList() {
	routeAuth := h.api.Group("/api/v1/tweets")
	routeAuth.Use(middleware.AuthMiddleware(secretKey))
	routeAuth.POST("/create", h.CreatePost)
}
