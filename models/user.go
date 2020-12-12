package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Username   string
	Password   string
	Todos   []Todo
}

// UserDTO model
type UserDTO struct {
	Username string   `form:"username"`
	Password string `form:"password"`
}

