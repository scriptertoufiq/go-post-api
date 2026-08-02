package user

import (
	"go-tweets/internal/service/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api         *gin.Engine
	validate    *validator.Validate
	userService user.UserService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, userService user.UserService) *Handler {
	return &Handler{
		api:         api,
		validate:    validate,
		userService: userService,
	}
}

func (h *Handler) RouteList() {
	authRoute := h.api.Group("/api/v1/auth")
	authRoute.POST("/register", h.RegisterUser)
}
