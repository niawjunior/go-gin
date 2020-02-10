package models

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	UsersData
}

type UsersData struct {
	Name string `form:"name" json:"name" sql:"not null" binding:"required"`
	Age  int    `json:"age"`
}
