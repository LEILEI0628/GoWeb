package ioc

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	type Config struct {
		DSN string `yaml:"dsn"` // ？此处DSN必须为大写
	}
	var cfg Config
	//var cfg Config = Config{DSN: "root:root@tcp(localhost:3306)/goweb"} // 设置默认值的方法2
	//dsn := viper.GetString("db.mysql.dsn")
	// 使用UnmarshalKey解析"db.mysql"时配置文件必须为树状结构：
	// db:
	//  mysql:
	//   dsn: ""
	// 不能写成db.mysql.dsn: ""的形式（使用配置文件和Remote方式时层级似乎不同，建议直接写成全树状结构）
	err := viper.UnmarshalKey("db.mysql", &cfg) // 需要对cfg进行修改，传指针
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(mysql.Open(cfg.DSN))
	if err != nil {
		// panic相当于整个goroutine结束
		// panic只会出现在初始化的过程中（一旦初始化出错，就没必要启动了）
		panic(err)
	}

	// 建表：
	//err = dao.InitUserTable(db)
	//if err != nil {
	//	panic(err)
	//}

	return db
}
