package models

import "gorm.io/gorm"

type BlogUser struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
