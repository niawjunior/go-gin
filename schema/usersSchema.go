package schema

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	UsersData
}

type UsersData struct {
	Name string `json:"name" sql:"not null" binding:"required"`
	Age  int    `json:"age"`
}
