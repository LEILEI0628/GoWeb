package ioc

import (
	"github.com/LEILEI0628/GoWeb/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
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
