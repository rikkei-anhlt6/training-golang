package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm: "column:username; not null; unique"`
	Age      int    `gorm: "column:age; not null; unique"`
}
