package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "12345678"
	DB_NAME     = "user1"
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	db = ConnectDB()
	return db
}
func ConnectDB() *gorm.DB {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Connect Success " + DB_NAME + " " + DB_HOST + ":" + DB_PORT)
	return db
}
