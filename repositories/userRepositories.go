package repositories

import (
	"gouserservice/config"
	"gouserservice/models"
)

func FindUserById(id string) (*models.User,error) {

	var user models.User

	result := config.DB.Where("id=?", id).First(&user)

	if result.Error != nil{

		return nil , result.Error
		//return models.User{}, result.Error
	}
	
	return &user, nil

}
