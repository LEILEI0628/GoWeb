package dao

import "gorm.io/gorm"

func InitUserTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
