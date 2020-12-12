package models

import "gorm.io/gorm"

// Todo model
type Todo struct {
	gorm.Model
	Name   string
	Body   string
	UserID uint
}

// TodoDTO model
type TodoDTO struct {
	Name string   `form:"name"`
	Body string `form:"body"`
}

