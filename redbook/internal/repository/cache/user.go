package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang-web-learn/redbook/internal/domain"
	"time"
)

var ErrKeyNotExist = redis.Nil

type UserCache struct {
	// A用到了B，B应当是接口；A用到了B，B应当是A的字段；A用到了B，A绝对不初始化B，而是通过外部注入
	// 传单机或cluster Redis都可以
	client     redis.Cmdable // 对外隐藏内部的实现
	expiration time.Duration
}

func NewUserCache(client redis.Cmdable, expiration time.Duration) *UserCache {
	// 不要在这里初始化！（传入配置或从系统获取都不要）
	return &UserCache{client: client, expiration: expiration}
}

func (userCache UserCache) Get(context context.Context, id int64) (domain.User, error) {
	value, err := userCache.client.Get(context, userCache.key(id)).Bytes()
	// 数据不存在，返回ErrKeyNotExist
	if err != nil {
		return domain.User{}, err
	}
	var user domain.User
	err = json.Unmarshal(value, &user)
	//if err != nil {
	//	return domain.User{}, err
	//}
	return user, nil // 此处为简写：有err则user一定为空
}

func (userCache UserCache) Set(context context.Context, user domain.User) error {
	value, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return userCache.client.Set(context, userCache.key(user.Id), value, userCache.expiration).Err()
}

func (userCache UserCache) key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}
