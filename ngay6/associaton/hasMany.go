package main

// User has many CreditCards, UserID is the foreign key
// Override Foreign Key, Override References, Polymorphism Association giống hasOne
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "12345678"
	DB_NAME     = "hasmany"
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
)

type User struct {
	gorm.Model
	CreditCard []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Connect Success " + DB_NAME + " " + DB_HOST + ":" + DB_PORT)
	db.AutoMigrate(&User{}, &CreditCard{})
}
