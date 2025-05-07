package routes

import (
	"gouserservice/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.GET("/", controllers.GetUser)
	router.POST("/", controllers.CreateUser)

}
