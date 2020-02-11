package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niawjunior/gin-app/database"
	"github.com/niawjunior/gin-app/models"
)

func GetAllUsers(c *gin.Context) {
	var users []models.Users
	db := database.GetDB()
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.Users
	db := database.GetDB()

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func AddUser(c *gin.Context) {
	var user models.Users
	var db = database.GetDB()

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		db.Create(&user)
		c.JSON(http.StatusCreated, &user)
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.Users
	var db = database.GetDB()
	c.BindJSON(&user)

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		db.Save(&user)
		c.JSON(http.StatusOK, &user)
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.Users
	db := database.GetDB()

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		db.Delete(&user)
	}
}
