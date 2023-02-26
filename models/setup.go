package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	//database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	dsn := "user:password@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Book{})

	if err != nil {
		return
	}

	DB = database
}
