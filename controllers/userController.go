package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	//c.String(http.StatusOK, "GET user inside user controller")
	c.JSON(http.StatusOK, "GET user inside user controller")
}

func CreateUser(c *gin.Context)  {
	c.String(http.StatusOK, "POST CreateUser")	
}