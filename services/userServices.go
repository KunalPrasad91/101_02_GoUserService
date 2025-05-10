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

func CreateUserService(dto models.UserRequestDto) (*models.UserResponseDto, error) {

	_, err := repositories.FindUserByEmail(dto.Email)

	// err == Nil means user with email found
	if err == nil {
		return nil, errors.New("User with email id already exist :- " + dto.Email)
	}

	user := models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	createduser, err := repositories.CreateUser(user)

	if err != nil {
		return nil, err
	}

	response := models.UserResponseDto{
		Id:    int(createduser.ID),
		Name:  createduser.Name,
		Email: createduser.Email,
	}

	return &response, nil

}

func DeleteUserService(id string) error {

	err := repositories.DeleteUserById(id)

	if err != nil {
		return errors.New("Failed to delete with user id :- " + id)
	}

	return nil

}

func UpdateUser(id string, dto models.UserRequestDto) (*models.UserResponseDto, error) {

	updates := make(map[string]interface{})

	if dto.Name != "" {
		updates["name"] = dto.Name
	}

	if dto.Email != "" {
		updates["email"] = dto.Email
	}

	if dto.Password != "" {
		updates["password"] = dto.Password
	}

	if len(updates) == 0 {
		return nil, errors.New("Nothing to update")
	}

	user, err := repositories.UpdateUser(id, updates)

	if err != nil {
		return nil, errors.New("Updates failed")
	}

	response := models.UserResponseDto{
		Id:    int(user.ID),
		Name:  user.Name,
		Email: user.Password,
	}

	return &response, nil
}
