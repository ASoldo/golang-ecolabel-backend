package services

import (
	"github.com/ASoldo/golang-ecolabel-backend/internal/config"
	"github.com/ASoldo/golang-ecolabel-backend/internal/errors"
	"github.com/ASoldo/golang-ecolabel-backend/internal/models"
)

// UserService is an interface that defines the contract for user-related services.
type UserService interface {
	Authenticate(username, password, groupID string) (string, error)
}

// UserServiceImpl is a struct that implements the UserService interface.
type UserServiceImpl struct{}

// NewUserService returns an instance of UserServiceImpl, which implements the UserService interface.
func NewUserService() UserService {
	return &UserServiceImpl{}
}

// Authenticate is a method of UserServiceImpl that checks if the provided username and password
// match the demo user's credentials. If they match, it returns a JWT token; otherwise, it returns an error.

func (s *UserServiceImpl) Authenticate(username, password, groupID string) (string, error) {
	if username == models.DemoUser.Username && password == models.DemoUser.Password &&
		groupID == models.DemoUser.GroupID {
		return models.GenerateToken(config.JwtSecret), nil
	}
	return "", errors.NewAppError(401, "Invalid credentials")
}
