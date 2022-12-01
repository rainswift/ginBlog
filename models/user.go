package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type BlogUser struct {
	gorm.Model
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
	Content   string `json:"content" binding:"required"`
	Title     string `json:"title" binding:"required"`
	ImgBg     string `json:"imgBg" binding:"required"`
	Label     string `json:"label" binding:"required"`
	Introduce string `json:"introduce" binding:"required"`
}

type GetId struct {
	Id int `json:"id"`
}

type UserInfo struct {
	UserId    int    `json:"userId"`
	Name      string `json:"name"`
	HeadImg   string `json:"headImg"`
	Introduce string `json:"introduce"`
	Github    string `json:"github"`
	Qq        string `json:"qq"`
}
