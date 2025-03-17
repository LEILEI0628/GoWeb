package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (userDAO *UserDAO) Insert(context context.Context, user User) error {
	now := time.Now().UnixMilli() // 当前时间的毫秒数
	user.CreateTime = now
	user.UpdateTime = now
	return userDAO.db.WithContext(context).Create(&user).Error
}

// User dao.User直接对应数据库表
// 其他叫法：entity，model，PO（persistent object）
type User struct {
	Id         int64  `gorm:"primaryKey,auto_Increment"` // 自增主键
	Email      string `gorm:"unique"`
	Password   string
	CreateTime int64 // 创建时间：毫秒数
	UpdateTime int64 // 修改时间：毫秒数

}
