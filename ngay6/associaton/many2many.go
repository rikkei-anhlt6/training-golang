package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Which creates join table: user_profiles
//   foreign key: user_refer_id, reference: users.refer
//   foreign key: profile_refer, reference: profiles.user_refer
type User1 struct {
	gorm.Model
	FullName  string
	Languages []Language `gorm:"many2many: user_languages;"`
	Profiles  []Profile  `gorm:"many2many: user_profiles; foreignkey: Refer; joinForeignkey: UserReferID; reference:ProfileRefer; joinReferences:ProfileRefer "`
	Refer     uint       `gorm:"index:,unique"`
}

type Language struct {
	gorm.Model
	LanguageName string
	// Users        []User1 `gorm:"many2many:user_languages; foreignkey: ID"` Back-Reference
}

type Profile struct {
	gorm.Model
	Name         string
	ProfileRefer uint `gorm:"index:,unique"`
}

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "12345678"
	DB_NAME     = "many2many"
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
	db.AutoMigrate(&User1{}, &Language{})
}
