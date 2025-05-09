package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
//	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
