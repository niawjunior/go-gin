package controllers

import (
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