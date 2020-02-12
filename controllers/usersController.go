package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niawjunior/go-gin/helpers"
	"github.com/niawjunior/go-gin/models"
	"github.com/niawjunior/go-gin/schema"
)

func GetAllUsers(c *gin.Context) {
	var users []schema.Users
	err := models.GetAllUsers(&users)
	if err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, users)
	} else {
		helpers.RespondJSON(c, http.StatusOK, users)
	}
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user schema.Users
	err := models.GetUserById(&user, id)
	if err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, user)
	} else {
		helpers.RespondJSON(c, http.StatusOK, user)
	}
}

func AddUser(c *gin.Context) {
	var user schema.Users
	c.BindJSON(&user)
	err := models.AddUser(&user)
	if err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, user)
	} else {
		helpers.RespondJSON(c, http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user schema.Users
	id := c.Param("id")
	err := models.GetUserById(&user, id)
	if err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = models.UpdateUser(&user, id)
	if err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, user)
	} else {
		helpers.RespondJSON(c, http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user schema.Users
	id := c.Param("id")
	err := models.DeleteUser(&user, id)
	if err != nil {
		helpers.RespondJSON(c, http.StatusNotFound, user)
	} else {
		helpers.RespondJSON(c, http.StatusOK, user)
	}
}
