package routes

import (
	"gouserservice/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	userGroup := router.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.POST("/", controllers.CreateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
		//router.PATCH("/:id", controllers.UpdateUser)
		userGroup.DELETE("/", controllers.DeleteAllUser)
		userGroup.PATCH("/:id", controllers.UpdateUser)
	}
}
