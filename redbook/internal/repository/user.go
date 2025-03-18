package repository

import (
	"context"
	"golang-web-learn/redbook/internal/domain"
	"golang-web-learn/redbook/internal/repository/dao"
)

var ErrUserEmailDuplicated = dao.ErrUserEmailDuplicated
var ErrUserNotFound = dao.ErrUserNotFound

type UserRepository struct {
	userDAO *dao.UserDAO
}

func NewUserRepository(userDAO *dao.UserDAO) *UserRepository {
	return &UserRepository{userDAO: userDAO}
}

func (userRepository *UserRepository) FindById(int64) {
	// 从cache中寻找

	// 从dao中寻找并写回cache
}

func (userRepository *UserRepository) Create(context context.Context, user domain.User) error {
	return userRepository.userDAO.Insert(context, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})

	// TODO 操作缓存
}

func (userRepository *UserRepository) FindByEmail(context context.Context, email string) (domain.User, error) {
	user, err := userRepository.userDAO.FindByEmail(context, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{Id: user.Id, Email: user.Email, Password: user.Password}, nil
}
