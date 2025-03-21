//go:build k8s

// 使用k8s编译标签
package config

var Config = config{
	DB: DBConfig{
		// 本地连接
		DSN: "root:20010628@tcp(redbook-mysql:13306)/redbook",
	},
	Redis: RedisConfig{
		Addr: "redbook-redis:16379",
	},
}
