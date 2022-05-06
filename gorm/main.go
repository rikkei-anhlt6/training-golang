package main

import (
	"fmt"
	"gorm-module/database"
	"gorm-module/models"

	"gorm.io/gorm"
)

var db *gorm.DB

func getAllUser() {
	var user models.User
	db.Find(&user)
	fmt.Println("Get all User: ")
	fmt.Println(user)
}
func createUser() {
	user := models.User{Username: "anhlta", Email: "anhlt6@rikkeisoft.com"}
	db.Create(&user)
}
func updateUser() {
	var user models.User
	id := 1
	db.First(&user, id) // tìm user có id = 1
	db.Model(&user).Updates(
		models.User{
			Username: "anhlt1",
			Email:    "anhlt1@rikkeisoft.com",
		})
	fmt.Println("User sau khi update", user)
	//fmt.Println(user)
}
func deleteUser() {
	var user models.User
	id := 2
	db.First(&user, id)
	if user.Username == "" {
		fmt.Println("Không tìm thấy User")
	} else {
		db.Delete(&user, id)
		fmt.Println("Delete thành công")
	}
}
func main() {
	db = database.InitDB() //Kết nối database
	db.AutoMigrate(&models.User{})
	getAllUser()
	createUser()
	updateUser()
	deleteUser()

}
