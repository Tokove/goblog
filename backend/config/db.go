package config

import (
	"goblog/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	// 获取要配置的信息
	dsn := AppConfig.Database.Dsn
	maxIdleConns := AppConfig.Database.MaxIdleConns
	maxOpenConns := AppConfig.Database.MaxOpenConns

	// 打开mysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open mysql: %v", err)
	}

	// 获取底层sql
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to configure database: %v", err)
	}

	// 设置配置信息
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 设置全局变量
	global.DB = db

	log.Println("Configure database successfully")
}
