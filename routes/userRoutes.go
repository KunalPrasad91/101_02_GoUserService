package routes

import (
	"gouserservice/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.GET("/", controllers.GetUsers)
	router.POST("/", controllers.CreateUser)
	router.DELETE("/:id", controllers.DeleteUser)
	//router.PATCH("/:id", controllers.UpdateUser)
	router.DELETE("/", controllers.DeleteAllUser)
}
