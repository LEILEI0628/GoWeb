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

type UserRepository interface {
	FindById(ctx context.Context, id int64) (domain.User, error)
	Create(ctx context.Context, user domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
}
type CacheUserRepository struct {
	userDAO   dao.UserDAO
	userCache cache.UserCache
}

func NewCacheUserRepository(dao *dao.GORMUserDAO, cache cache.UserCache) UserRepository {
	return &CacheUserRepository{userDAO: dao, userCache: cache}
}

func (repo *CacheUserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	// 从cache中寻找
	// 获取用户缓存
	uc, err := repo.userCache.Get(ctx, id)
	if err == nil {
		// 从cache中找到数据
		fmt.Println("Cache Find")
		return uc, err
	}
	//if errors.Is(err, cachex.ErrKeyNotExist) { // 处理缓存未命中：从cache中没找到数据
	// 设置用户缓存：从dao中寻找并写回cache
	ue, err := repo.userDAO.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	ud := repo.entityToDomain(ue)
	err = repo.userCache.Set(ctx, ud.Id, ud)
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

func (repo *CacheUserRepository) Create(ctx context.Context, user domain.User) error {
	ue := repo.domainToEntity(user)
	return repo.userDAO.Insert(ctx, ue)

	// TODO 操作缓存
}

func (repo *CacheUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := repo.userDAO.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.entityToDomain(user), nil
}

func (repo *CacheUserRepository) domainToEntity(u domain.User) dao.User {
	return dao.User{Id: u.Id, Email: sql.NullString{String: u.Email, Valid: u.Email != ""}, Phone: sql.NullString{String: u.Phone, Valid: u.Phone != ""}, Password: u.Password, CreateTime: u.CreateTime.UnixMilli()}
}

func (repo *CacheUserRepository) entityToDomain(u dao.User) domain.User {
	return domain.User{Id: u.Id, Email: u.Email.String, Phone: u.Phone.String, Password: u.Password, CreateTime: time.UnixMilli(u.CreateTime)}
}
