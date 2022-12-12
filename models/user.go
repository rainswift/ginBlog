package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type SetTime struct {
	//ID        uint `gorm:"primarykey"`
	createda_at Time
	updated_at  Time
	//DeletedAt gorm.DeletedAt `gorm:"index"`
}
type BlogUser struct {
	gorm.Model
	SetTime
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}
type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

type Content struct {
	gorm.Model
	UserId    int    `json:"userId"`
	Content   string `json:"content" binding:"required"`
	Title     string `json:"title" binding:"required"`
	ImgBg     string `json:"imgBg" binding:"required"`
	Label     string `json:"label" binding:"required"`
	Introduce string `json:"introduce" binding:"required"`
}

type GetId struct {
	Id int `json:"id" form:"id" `
}

type UserInfo struct {
	UserId    int    `json:"userId"`
	Name      string `json:"name"`
	HeadImg   string `json:"headImg"`
	Introduce string `json:"introduce"`
	Github    string `json:"github"`
	Qq        string `json:"qq"`
}
