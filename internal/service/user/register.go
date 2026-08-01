package user

import (
	"context"
	"errors"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Register(ctx context.Context, req *dto.RegisterUserRequest) (int64, int, error) {
	// Implement the logic for registering a user here
	// check user already exists in the database
	userExists, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, req.Username)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	if userExists != nil {
		return 0, http.StatusBadRequest, errors.New("user already exists")
	}

	// Hashpass
	password, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	// Create user in the database
	now := time.Now()
	user := &model.UserModel{
		Name:      req.Name,
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(password),
		CreatedAt: now,
		UpdatedAt: now,
	}
	userID, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return userID, http.StatusCreated, nil
}
