package models

import "gorm.io/gorm"

type BlogUser struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
