package service

import (
	"context"
	"errors"
	"golang-web-learn/redbook/internal/domain"
	"golang-web-learn/redbook/internal/repository"
	"golang-web-learn/redbook/internal/repository/dao"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserEmailDuplicated = repository.ErrUserEmailDuplicated
var ErrInvalidEmailOrPassword = errors.New("invalid email or password")

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

func (userService *UserService) SignIn(context context.Context, user domain.User) (domain.User, error) {
	// 通过邮箱查找用户
	userFind, err := userService.userRepository.FindByEmail(context, user.Email)
	if errors.Is(err, dao.ErrUserNotFound) {
		return domain.User{}, ErrInvalidEmailOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	// 比较密码
	err = bcrypt.CompareHashAndPassword([]byte(userFind.Password), []byte(user.Password))
	if err != nil {
		return domain.User{}, ErrInvalidEmailOrPassword
	}
	return userFind, nil

}
