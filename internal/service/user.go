package service

import (
	"context"
	"errors"
	"github.com/LEILEI0628/GoWeb/internal/domain"
	"github.com/LEILEI0628/GoWeb/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserEmailDuplicated = repository.ErrUserEmailDuplicated
var ErrInvalidEmailOrPassword = errors.New("invalid email or password")
var ErrUserNotFound = repository.ErrUserNotFound

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
	if errors.Is(err, repository.ErrUserNotFound) {
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

func (userService *UserService) Profile(context context.Context, id int64) (domain.UserProfile, error) {
	user, err := userService.userRepository.FindById(context, id)
	if err != nil {
		return domain.UserProfile{}, err
	}
	return domain.UserProfile{Email: user.Email}, err
}
