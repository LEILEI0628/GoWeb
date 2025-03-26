package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/LEILEI0628/GinPro/middleware/cache"
	"github.com/LEILEI0628/GoWeb/internal/domain"
	"github.com/LEILEI0628/GoWeb/internal/repository/cache"
	"github.com/LEILEI0628/GoWeb/internal/repository/dao"
	"time"
)

var ErrUserEmailDuplicated = dao.ErrUserEmailDuplicated
var ErrUserNotFound = dao.ErrUserNotFound
var ErrKeyNotExist = cachex.ErrKeyNotExist

type UserRepository struct {
	userDAO   *dao.UserDAO
	userCache *cache.UserCache
}

func NewUserRepository(userDAO *dao.UserDAO, userCache *cache.UserCache) *UserRepository {
	return &UserRepository{userDAO: userDAO, userCache: userCache}
}

func (userRepository *UserRepository) FindById(context context.Context, id int64) (domain.User, error) {
	// 从cache中寻找
	// 获取用户缓存
	uc, err := userRepository.userCache.Get(context, id)
	if err == nil {
		// 从cache中找到数据
		fmt.Println("Cache Find")
		return uc, err
	}
	//if errors.Is(err, cachex.ErrKeyNotExist) { // 处理缓存未命中：从cache中没找到数据
	// 设置用户缓存：从dao中寻找并写回cache
	ue, err := userRepository.userDAO.FindById(context, id)
	if err != nil {
		return domain.User{}, err
	}
	ud := userRepository.entityToDomain(ue)
	err = userRepository.userCache.Set(context, ud.Id, ud)
	if err != nil {
		fmt.Println("Cache Set Filed")
		// 缓存Set失败（记录日志做监控即可，为了防止缓存崩溃的可能）
		// TODO 记录日志
	}
	fmt.Println("Cache Set Success")
	return ud, err
	//} // 注释掉此处if语句代表不管缓存发生什么问题都从数据库加载
	// 当缓存发生除ErrKeyNotExist的错误时由两种解决方案：
	// 1.从数据库加载（偶发错误友好），极个别缓存错误可以解决，但当缓存真的崩溃时，要做好兜底保护数据库（大量访问）
	// 2.不加载，默认缓存崩溃，极个别缓存错误也不解决，用户体验较差
	// 面试时选1，极个别缓存错误可以解决，缓存真的崩溃时可以选择数据库限流（基于内存的单机限流）、布尔过滤器
}

func (userRepository *UserRepository) Create(context context.Context, user domain.User) error {
	ue := userRepository.domainToEntity(user)
	return userRepository.userDAO.Insert(context, ue)

	// TODO 操作缓存
}

func (userRepository *UserRepository) FindByEmail(context context.Context, email string) (domain.User, error) {
	user, err := userRepository.userDAO.FindByEmail(context, email)
	if err != nil {
		return domain.User{}, err
	}
	return userRepository.entityToDomain(user), nil
}

func (userRepository *UserRepository) domainToEntity(u domain.User) dao.User {
	return dao.User{Id: u.Id, Email: sql.NullString{String: u.Email, Valid: u.Email != ""}, Phone: sql.NullString{String: u.Phone, Valid: u.Phone != ""}, Password: u.Password, CreateTime: u.CreateTime.UnixMilli()}
}

func (userRepository *UserRepository) entityToDomain(u dao.User) domain.User {
	return domain.User{Id: u.Id, Email: u.Email.String, Phone: u.Phone.String, Password: u.Password, CreateTime: time.UnixMilli(u.CreateTime)}
}
