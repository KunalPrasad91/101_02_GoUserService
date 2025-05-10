package repositories

import (
	"errors"
	"gouserservice/config"
	"gouserservice/models"
)

func FindUserById(id string) (*models.User, error) {

	var user models.User

	result := config.DB.Where("id=?", id).First(&user)

	if result.Error != nil {

		return nil, result.Error
		//return models.User{}, result.Error
	}

	return &user, nil

}

func FindUserByEmail(email string) (*models.User, error) {

	var user models.User

	result := config.DB.Where("email=?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}

func CreateUser(user models.User) (*models.User, error) {

	result := config.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}

func DeleteUserById(id string) error {

	var user models.User
	result := config.DB.Where("id=?", id).First(&user)

	if result.Error != nil {
		return result.Error
	}

	err := config.DB.Delete(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(id string, updates map[string]interface{}) (*models.User, error) {

	var user models.User
	result := config.DB.Where("id=?", id).First(&user)

	if result.Error != nil {
		return nil, errors.New("User with id dont exist :- " + id )
	}

	err := config.DB.Model(&user).Updates(updates)

	if err!=nil{
		return nil,  errors.New("User with id Update failed :- " + id )
	}

	return &user, nil
}
