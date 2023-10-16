package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	mod "picme-backend/modules"
	"picme-backend/types/table"
)

func Init() {
	dsn := "poc:poc1212312121@tcp(server2.bsthun.com:4004)/poc1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect database")
	}

	if err := db.AutoMigrate(new(table.Users)); err != nil {
		panic(err)
	}

	mod.DB = db
}
