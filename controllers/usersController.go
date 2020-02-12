package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/niawjunior/gin-app/helpers"
	"github.com/niawjunior/gin-app/models"
	"github.com/niawjunior/gin-app/schema"
)

func GetAllUsers(c *gin.Context) {
	var users []schema.Users
	err := models.GetAllUsers(&users)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, users)
	} else {
		ApiHelpers.RespondJSON(c, 200, users)
	}
}

func GetUserById(user *schema.Users, id string) (err error) {
	id := c.Param("id")
	var user schema.Users
	err := models.GetUserById(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func AddUser(c *gin.Context) {
	var user schema.Users
	c.BindJSON(&user)
	err := models.AddUser(&user)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user schema.Users
	id := c.Param("id")
	err := models.GetUserById((&user, id))
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	}
	c.BindJSON(&user)
	err = models.UpdateUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user Models.Book
	id := c.Param("id")
	err := models.DeleteUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}