package services

import (
	"gouserservice/models"
	"gouserservice/repositories"

	"errors"
)

func GetUserbyId(id string) (*models.UserResponseDto, error) {

	user, err := repositories.FindUserById(id)

	if err != nil {
		return nil, errors.New("User Not Found with id :- " + id)
	}

	return &models.UserResponseDto{
		Id:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil

}
