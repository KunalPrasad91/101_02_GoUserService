package controllers

import (
	"gouserservice/config"
	"gouserservice/models"
	"gouserservice/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var dto models.UserRequestDto
	err := c.BindJSON(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := services.CreateUserService(dto)

	if err != nil {

		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)

}

func GetUsers(c *gin.Context) {

	var users []models.User

	config.DB.Find(&users)

	c.JSON(http.StatusOK, &users)

}

func GetUser(c *gin.Context) {

	id := c.Param("id")

	user, err := services.GetUserbyId(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	err := services.DeleteUserService(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted Successfully id :-" + id})

}

func DeleteAllUser(c *gin.Context) {

	//result := config.DB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	result := config.DB.Exec("TRUNCATE TABLE users")

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All users deleted successfully"})

}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")

	var dto models.UserRequestDto

	err := c.BindJSON(&dto)

	if err!=nil{
		c.JSON(http.StatusBadRequest, err.Error())
	}
	
	createduser, err := services.UpdateUser(id,dto)

	if err!=nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "User updated successully",
								"user" : createduser})

}
