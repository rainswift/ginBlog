package dao

import (
	"fmt"
	"ginBlog/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *models.BlogUser)
	GetByName(name string) bool
	GetLoadUser(name string) (*models.BlogUser, bool)
	Cs_text()
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&models.BlogUser{})
}

//func (mgr *manager) AddUser(user *models.BlogUser) {
//	fmt.Printf("% + v\n", user)
//	mgr.db.Create(user)
//}

// 创建用户
func (mgr *manager) AddUser(user *models.BlogUser) {
	mgr.db.Create(user)

}

// 根据用户名查询用户
func (mgr *manager) GetByName(name string) bool {
	var user models.BlogUser

	//mgr.db.Last(&user)
	//result := mgr.db.Where("username = ?", "qq")
	//affected := result.RowsAffected
	//fmt.Printf("% + v\n",result)
	//fmt.Println(affected)

	result := mgr.db.Where("username = ?", name).First(&user, "username = ?", name)
	//if err != nil {
	//	return false
	//}
	affected := result.RowsAffected
	if affected >= 1 {
		return false
	} else {
		return true
	}
}
func (mgr *manager) GetLoadUser(name string) (*models.BlogUser, bool) {
	var u models.BlogUser
	var flag bool
	result := mgr.db.Where("username=?", name).First(&u)
	affected := result.RowsAffected
	if affected >= 1 {
		flag = false
	} else {
		flag = true
	}
	return &u, flag
}

func (mgr *manager) Cs_text() {
	var user models.BlogUser
	result := mgr.db.Last(&user)
	affected := result.RowsAffected
	fmt.Printf("% + v\n", user.Username)
	fmt.Println(affected)
}
