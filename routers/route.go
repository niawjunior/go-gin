package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/niawjunior/gin-app/controllers"
)

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", controllers.GetAllUsers)
			users.POST("/", controllers.AddUser)
			users.PUT("/:id", controllers.UpdateUser)
		}
	}
	return r
}
