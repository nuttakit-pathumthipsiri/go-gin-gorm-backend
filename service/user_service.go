package service

import (
	"errors"
	"go-gin-gorm-backend/config"
	"go-gin-gorm-backend/model"
	"go-gin-gorm-backend/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// LoginUser authenticates a user and returns tokens
func (s *UserService) LoginUser(req *model.LoginRequest) (*model.LoginResponse, error) {
	// Get user by username
	user, err := s.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Check password
	if !config.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	// Generate tokens
	token, err := config.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	// Create response
	response := &model.LoginResponse{
		Token: token,
		User:  user.ToUserResponse(),
	}

	return response, nil
}
