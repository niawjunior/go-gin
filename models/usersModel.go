package models

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	UsersData
}

type UsersData struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age string `json:"age"`
}