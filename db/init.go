package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// Init init DB
func Init() {
	zap.S().Info("=============>初始化MYSQL数据库连接<==============")
	var err error
	DB, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/classdesign?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
		zap.S().Error("=============>数据库连接失败<==============")
	}
}
