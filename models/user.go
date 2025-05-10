package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserRequestDto struct {
	Name 		string `json:"name"`
	Email		string `json:"password"`
	Password	string `json:"email"`
}

type UserResponseDto struct {
	Id          int    `json:"id"`
	Name 		string `json:"name"`
	Email		string `json:"password"`
}
