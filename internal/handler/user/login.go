package user

import (
	"net/http"

	"go-tweets/internal/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, refreshToken, statusCode, err := h.userService.Login(ctx, &req)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(statusCode, dto.LoginUserResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}
