package service

import (
	"github.com/gin-gonic/gin"
	"golang-web-learn/redbook/internal/domain"
)

type UserService struct{}

func (userService *UserService) SignUp(context *gin.Context, user domain.User) error {
	return nil
}
