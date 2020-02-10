package models

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	UsersData
}

type UsersData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
