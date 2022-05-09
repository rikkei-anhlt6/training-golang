package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mỗi người dùng chỉ có thể có 1 CreditCard
// CreditCardID is the foreign key
type User struct {
	gorm.Model
	Name       string     `gorm:"index"`
	CreditCard CreditCard `gorm:"foreignkey: UserName; references: name; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` //Override Foreign Key, thay UserID thành UserName làm foreignkey
	// references dùng để tham chiếu đến trường
	// Bang CreditCard chứa khoá ngoại tên là username tham chiếu đến trường name của bảng User
}
type CreditCard struct {
	gorm.Model
	Number   string
	UserID   uint
	UserName string `gorm:"size:191"`
}
type Cat struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "12345678"
	DB_NAME     = "hasone"
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
)

func main() {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Connect Success " + DB_NAME + " " + DB_HOST + ":" + DB_PORT)
	db.AutoMigrate(&User{}, &CreditCard{}, &Dog{}, &Toy{}, &Cat{})
	db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})
	// INSERT INTO `dogs` (`name`) VALUES ("dog1")
	// INSERT INTO `toys` (`name`,`owner_id`,`owner_type`) VALUES ("toy1","1","dogs"), id của dog1 là 1 => owner_id = 1
}
