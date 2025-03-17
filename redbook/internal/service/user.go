package service

import (
	"context"
	"golang-web-learn/redbook/internal/domain"
	"golang-web-learn/redbook/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (userService *UserService) SignUp(context context.Context, user domain.User) error {
	return userService.userRepository.Create(context, user)
}
