package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserEmailDuplicated = errors.New("user email duplicated")
	ErrUserNotFound        = gorm.ErrRecordNotFound // 别名，返回的gorm.ErrRecordNotFound err会被改写成ErrUserNotFound
)

type UserDAO interface {
	Insert(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (User, error)
	FindById(ctx context.Context, id int64) (User, error)
}
type GORMUserDAO struct {
	db *gorm.DB
}

func NewGORMUserDAO(db *gorm.DB) *GORMUserDAO {
	return &GORMUserDAO{db: db}
}

// User dao.User直接对应数据库表
// 其他叫法：entity，model，PO（persistent object）
type User struct {
	Id         int64          `gorm:"primaryKey,auto_Increment"` // 自增主键
	Email      sql.NullString `gorm:"unique"`                    // 唯一索引允许有多个空值（null），但是不能有多个空字符串（""）
	Phone      sql.NullString `gorm:"unique"`                    // *string也可以，但是要解引用，判空
	Password   string
	CreateTime int64 // 创建时间：毫秒数
	UpdateTime int64 // 修改时间：毫秒数

}

func (dao *GORMUserDAO) Insert(ctx context.Context, user User) error {
	now := time.Now().UnixMilli() // 当前时间的毫秒数
	user.CreateTime = now
	user.UpdateTime = now
	err := dao.db.WithContext(ctx).Create(&user).Error
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

func (dao *GORMUserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	// SELECT * FROM `users` WHERE `email`=?
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&user).Error
	//err := dao.db.WithContext(context).First(&user, "email=?", email).Error // 等价写法
	return user, err // 无需判断直接返回（已更改err名称）
}

func (dao *GORMUserDAO) FindById(ctx context.Context, id int64) (User, error) {
	var user User
	// SELECT * FROM `users` WHERE `id`=?
	err := dao.db.WithContext(ctx).Where("id=?", id).First(&user).Error
	return user, err // 无需判断直接返回（已更改err名称）
}
