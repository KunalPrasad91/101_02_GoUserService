package controllers

import (
	"gouserservice/config"
	"gouserservice/models"
	"gouserservice/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user models.User

	c.BindJSON(&user)

	config.DB.Create(&user)

	c.JSON(http.StatusOK, &user)
}

func GetUsers(c *gin.Context) {

	var users []models.User

	config.DB.Find(&users)

	c.JSON(http.StatusOK, &users)

}

func GetUser(c *gin.Context)  {
	
	id := c.Param("id")

	user, err := services.GetUserbyId(id)

	if err!=nil{

		c.JSON(http.StatusNotFound, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {

	var user models.User
	config.DB.Where("id=?", c.Param("id")).Delete(&user)

	c.JSON(http.StatusOK, &user)
}

func DeleteAllUser(c *gin.Context)  {

	
	//result := config.DB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	result := config.DB.Exec("TRUNCATE TABLE users")
	
	if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "All users deleted successfully"})

}
