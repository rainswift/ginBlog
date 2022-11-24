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

type Content struct {
	gorm.Model
	Content   string `json:"content" binding:"required"`
	Title     string `json:"title" binding:"required"`
	ImgBg     string `json:"imgBg" binding:"required"`
	Label     string `json:"label" binding:"required"`
	Introduce string `json:"introduce" binding:"required"`
}

type GetId struct {
	Id int `json:"id"`
}
