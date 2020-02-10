package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/niawjunior/go-app/controllers"
	db "github.com/niawjunior/go-app/database"
)

func main() {
	db.Init()
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", c.GetAllUsers)
		}
	}
	r.Run(":3001")
}
