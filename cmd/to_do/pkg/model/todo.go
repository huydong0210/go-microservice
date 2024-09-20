package model

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Name   string
	State  string
	UserId uint
}
