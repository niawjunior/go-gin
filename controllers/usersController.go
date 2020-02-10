package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niawjunior/go-app/database"
	"github.com/niawjunior/go-app/models"
)

func GetAllUsers(c *gin.Context) {
	var users []models.Users
	db := database.GetDB()
	db.Find(&users)
	c.JSON(200, users)
}

func AddUser(c *gin.Context) {
	var user models.Users
	var db = database.GetDB()

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		db.Create(&user)
		c.JSON(http.StatusOK, &user)
	}
}
