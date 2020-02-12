package models

import (
	"github.com/niawjunior/go-gin/database"
	"github.com/niawjunior/go-gin/schema"
)

func GetAllUsers(users *[]schema.Users) (err error) {
	if err = database.DB.Find(users).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(user *schema.Users, id string) (err error) {
	if err := database.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func AddUser(user *schema.Users) (err error) {
	if err = database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *schema.Users, id string) (err error) {
	database.DB.Save(user)
	return nil
}

func DeleteUser(user *schema.Users, id string) (err error) {
	database.DB.Where("id = ?", id).Delete(user)
	return nil
}
