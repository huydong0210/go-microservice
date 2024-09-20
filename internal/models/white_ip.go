package models

import "gorm.io/gorm"

type WhiteIP struct {
	gorm.Model
	Ip          string `json:"ip"`
	Description string `json:"description"`
}
