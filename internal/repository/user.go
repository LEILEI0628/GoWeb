package repository

import (
	"context"
	"github.com/LEILEI0628/GoWeb/internal/domain"
	"github.com/LEILEI0628/GoWeb/internal/repository/cache"
	"github.com/LEILEI0628/GoWeb/internal/repository/dao"
)

var ErrUserEmailDuplicated = dao.ErrUserEmailDuplicated
var ErrUserNotFound = dao.ErrUserNotFound
var ErrKeyNotExist = cache.ErrKeyNotExist

type UserRepository struct {
	userDAO   *dao.UserDAO
	userCache *cache.UserCache
}

func NewUserRepository(userDAO *dao.UserDAO, userCache *cache.UserCache) *UserRepository {
	return &UserRepository{userDAO: userDAO, userCache: userCache}
}

func (userRepository *UserRepository) FindById(context context.Context, id int64) (domain.User, error) {
	// 从cache中寻找
	user, err := userRepository.userCache.Get(context, id)
	if err == nil { // 从cache中找到数据
		return user, err
	}
	//if errors.Is(err, ErrKeyNotExist) { // 从cache中没找到数据
	// 从dao中寻找并写回cache
	userEntity, err := userRepository.userDAO.FindById(context, id)
	if err != nil {
		return domain.User{}, err
	}
	user = domain.User{Id: userEntity.Id, Email: userEntity.Email, Password: userEntity.Password}
	err = userRepository.userCache.Set(context, user)
	if err != nil {
		// 缓存Set失败（记录日志做监控即可，为了防止缓存崩溃的可能）
		// TODO 记录日志
	}
	return user, err
	//} // 注释掉此处if语句代表不管缓存发生什么问题都从数据库加载
	// 当缓存发生除ErrKeyNotExist的错误时由两种解决方案：
	// 1.从数据库加载（偶发错误友好），极个别缓存错误可以解决，但当缓存真的崩溃时，要做好兜底保护数据库（大量访问）
	// 2.不加载，默认缓存崩溃，极个别缓存错误也不解决，用户体验较差
	// 面试时选1，极个别缓存错误可以解决，缓存真的崩溃时可以选择数据库限流（基于内存的单机限流）、布尔过滤器
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
