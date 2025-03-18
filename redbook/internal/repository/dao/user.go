package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserEmailDuplicated = errors.New("user email duplicated")
	ErrUserNotFound        = gorm.ErrRecordNotFound // 别名，返回的gorm.ErrRecordNotFound err会被改写成ErrUserNotFound
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
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

func (userDAO *UserDAO) Insert(context context.Context, user User) error {
	now := time.Now().UnixMilli() // 当前时间的毫秒数
	user.CreateTime = now
	user.UpdateTime = now
	err := userDAO.db.WithContext(context).Create(&user).Error
	// 关于邮箱冲突问题：使用先查操作（查询时都返回查询邮箱不存在）会导致后插入的用户操作失败
	// 尝试锁住（间隙锁）：SELECT * FROM users WHERE email=... FOR UPDATE（存在并发问题）
	// 冲突发生概率不大时可以不用分布式锁
	// 以下代码与底层强耦合（更换数据库失效）
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			return ErrUserEmailDuplicated
		}
	}
	return err
}

func (userDAO *UserDAO) FindByEmail(context context.Context, email string) (User, error) {
	var user User
	// SELECT * FROM `users` WHERE `email`=?
	err := userDAO.db.WithContext(context).Where("email=?", email).First(&user).Error
	//err := userDAO.db.WithContext(context).First(&user, "email=?", email).Error // 等价写法
	return user, err // 无需判断直接返回（已更改err名称）
}
