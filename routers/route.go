package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/niawjunior/go-gin/controllers"
)

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", controllers.GetAllUsers)
			users.GET("/:id", controllers.GetUserById)
			users.POST("/", controllers.AddUser)
			users.PUT("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}
	return r
}
