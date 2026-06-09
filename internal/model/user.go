package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" validate:"required,username_len"`
	Email string `json:"email" validate:"required,email"`
}
