package service

import (
	"context"
	"golang-web-learn/redbook/internal/domain"
	"golang-web-learn/redbook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserEmailDuplicated = repository.ErrUserEmailDuplicated

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (userService *UserService) SignUp(context context.Context, user domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return userService.userRepository.Create(context, user)
}
