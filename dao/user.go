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
	GetUserList(page *models.Pagination) []models.BlogUser
	UserDelete(id int) models.BlogUser
	SaveEdit(c *models.Content)
	GetEditList(page *models.Pagination) ([]models.Content, int64)
	GetEditDetails(id int) models.Content
	UserSave(c *models.UserInfo)
	GetUserInfo() models.UserInfo
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
	db.AutoMigrate(&models.Content{})
	db.AutoMigrate(&models.UserInfo{})
}

// 创建用户
func (mgr *manager) AddUser(user *models.BlogUser) {
	mgr.db.Create(user)
}

// 保存文章
func (mgr *manager) SaveEdit(c *models.Content) {
	mgr.db.Create(c)
}

// 保持个人信息
func (mgr *manager) UserSave(c *models.UserInfo) {
	mgr.db.Create(c)
}

// 根据用户名查询用户是否存在
func (mgr *manager) GetByName(name string) bool {
	var user models.BlogUser
	result := mgr.db.Where("username = ?", name).First(&user, "username = ?", name)
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

// 查找用户列表
func (mgr *manager) GetUserList(page *models.Pagination) []models.BlogUser {
	var users []models.BlogUser
	offset := (page.Page - 1) * page.Limit
	queryBuider := mgr.db.Limit(page.Limit).Offset(offset).Find(&users)
	queryBuider.Model(&models.BlogUser{}).Where(
		"username like ?  ",
		"%"+page.Sort+"%",
	).Find(&users)
	//mgr.db.Scopes(paginate(categories, &pagination, mgr.db)).Find(&categories)
	return users
}

// 用户删除
func (mgr *manager) UserDelete(id int) models.BlogUser {
	var users models.BlogUser
	mgr.db.Delete(&models.BlogUser{}, id).Find(&users)
	return users
}

// 查找文章列表
func (mgr *manager) GetEditList(page *models.Pagination) ([]models.Content, int64) {
	var content []models.Content
	len := mgr.db.Find(&content).RowsAffected
	fmt.Print(len)
	offset := (page.Page - 1) * page.Limit
	queryBuider := mgr.db.Limit(page.Limit).Offset(offset).Find(&content)
	queryBuider.Model(&models.BlogUser{}).Find(&content)

	//mgr.db.Scopes(paginate(categories, &pagination, mgr.db)).Find(&categories)
	return content, len
}

// 查找文章详情
func (mgr *manager) GetEditDetails(id int) models.Content {
	var content models.Content
	mgr.db.Where("id=?", id).First(&content)

	//mgr.db.Scopes(paginate(categories, &pagination, mgr.db)).Find(&categories)
	return content
}

// 查找用户信息
func (mgr *manager) GetUserInfo() models.UserInfo {
	var content models.UserInfo
	mgr.db.First(&content)
	//mgr.db.Scopes(paginate(categories, &pagination, mgr.db)).Find(&categories)
	return content
}
