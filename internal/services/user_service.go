package services

import (
	"github.com/ASoldo/golang-ecolabel-backend/internal/config"
	"github.com/ASoldo/golang-ecolabel-backend/internal/errors"
	"github.com/ASoldo/golang-ecolabel-backend/internal/models"
)

type UserService interface {
	Authenticate(username, password string) (string, error)
}

type UserServiceImpl struct{}

func NewUserService() UserService {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) Authenticate(username, password string) (string, error) {
	if username == models.DemoUser.Username && password == models.DemoUser.Password {
		return models.GenerateToken(config.JwtSecret), nil
	}
	return "", errors.NewAppError(401, "Invalid credentials")
}
