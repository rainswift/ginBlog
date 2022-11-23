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
	Content string `json:"content"`
	Title   string `json:"title"`
	ImgBg   string `json:"imgBg"`
	Label   string `json:"label"`
}
