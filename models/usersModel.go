package models

import (
	"github.com/niawjunior/go-gin/database"
	"github.com/niawjunior/go-gin/schema"
)

func GetAllUsers(users *[]schema.Users) (err error) {
	db := database.GetDB()
	if err = db.Find(users).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(user *schema.Users, id string) (err error) {
	db := database.GetDB()
	if err := db.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func AddUser(user *schema.Users) (err error) {
	db := database.GetDB()
	if err = db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *schema.Users, id string) (err error) {
	db := database.GetDB()
	db.Save(user)
	return nil
}

func DeleteUser(user *schema.Users, id string) (err error) {
	db := database.GetDB()
	db.Where("id = ?", id).Delete(user)
	return nil
}
