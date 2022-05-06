package models

import (
	"time"
)

type User struct {
	Id        uint32    `gorm: "column:id; primary_key; auto_increment"`
	Username  string    `gorm: "column:username; not null; unique"`
	Email     string    `gorm: "column:email; not null; unique"`
	CreatedAt time.Time `gorm: "column: CreatedAt; default:CURRENT_TIMESTAMP"`
	UpdateAt  time.Time `gorm: "column: UpdatedAt; default:CURRENT_TIMESTAMP"`
}
