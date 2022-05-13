package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"socket/server/database"
	"socket/server/models"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = database.InitDB()
	server, err := net.Listen("tcp", ":3003")
	db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		for {
			reader := bufio.NewReader(conn)
			msg, err1 := reader.ReadString('\n')
			fmt.Print(msg)
			if err1 != nil {
				break
			}
			user := models.User{}
			json.Unmarshal([]byte(msg), &user)
			db.Create(&user)

		}
	}
}
