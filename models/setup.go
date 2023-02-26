package models

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

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

func ConnectCache() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       0,
	})

	// pong, err := client.Ping().Result()
	// fmt.Println(pong, err)

	RDB = client
}
