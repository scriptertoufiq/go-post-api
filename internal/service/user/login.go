package user

import (
	"context"
	"fmt"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"go-tweets/pkg/jwt"
	"go-tweets/pkg/refreshtoken"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(ctx context.Context, req *dto.LoginUserRequest) (string, string, int, error) {
	//user is registered or not
	userExists, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, "")
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	if userExists == nil {
		return "", "", http.StatusNotFound, fmt.Errorf("wrong email")
	}

	// generate access token
	pass := bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(req.Password))
	if pass != nil {
		return "", "", http.StatusUnauthorized, fmt.Errorf("wrong email or password")
	}

	//get refresh token if exists otherwise generate new refresh token
	token, err := jwt.CreateToken(int64(userExists.ID), userExists.Username, s.cfg.SecretKey)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// generate & store refresh token
	now := time.Now()
	refreshTokenExists, err := s.userRepo.GetRefreshToken(ctx, int64(userExists.ID), now)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenExists != nil {
		return token, refreshTokenExists.Token, http.StatusOK, nil
	}

	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	err = s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:    int64(userExists.ID),
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour), // Set the expiration time for the refresh token (e.g., 7 days)
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	return token, refreshToken, http.StatusOK, nil

}
