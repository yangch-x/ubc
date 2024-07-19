package models

import (
	"gorm.io/gorm"
)

var (
	mysqlDb *gorm.DB
)

func InitModel2(mysql *gorm.DB) error {
	mysqlDb = mysql
	mysqlDb = mysqlDb.Debug()
	return nil
}

func InitModel(mysql *gorm.DB) error {
	mysqlDb = mysql
	mysqlDb = mysqlDb.Debug()
	// 自动同步更新表结构
	return mysqlDb.AutoMigrate()
}
