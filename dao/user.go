package dao

import (
	"fmt"
	"ginBlog/config"
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
	SaveEdit(c *models.Content, id string)
	GetEditList(page *models.Pagination, userId int) ([]models.Content, int64)
	GetEditDetails(id string, userId int) models.Content
	UserSave(c *models.UserInfo)
	GetUserInfo(id int) models.UserInfo
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

var AppConfig = &config.Configuration{}

func init() {
	cfgFile := "./config/config.yaml"
	conf, err := config.GetAllConfigValues(cfgFile)
	AppConfig = conf

	dsn := AppConfig.DatabaseSettings.DatabaseURI
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
func (mgr *manager) SaveEdit(c *models.Content, id string) {
	fmt.Println(c.ID)
	mgr.db.Model(c).Where("id = ?", c.ID).Save(c)
}

// 保持个人信息
func (mgr *manager) UserSave(c *models.UserInfo) {
	context := c
	mgr.db.Model(c).Where("user_id = ?", c.UserId).Save(context)

}

// 根据用户id查询用户信息是否存在
func (mgr *manager) GetByName(name string) bool {
	var user models.UserInfo
	result := mgr.db.Where("user_id = ?", name).First(&user, "username = ?", name)
	affected := result.RowsAffected
	if affected >= 1 {
		return false
	} else {
		return true
	}
}

// 根据用户名查询用户是否存在
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
func (mgr *manager) GetEditList(page *models.Pagination, userId int) ([]models.Content, int64) {
	var content []models.Content
	len := mgr.db.Where("user_id = ?", userId).Find(&content).RowsAffected
	fmt.Print(len)
	offset := (page.Page - 1) * page.Limit
	queryBuider := mgr.db.Where("user_id = ?", userId).Limit(page.Limit).Offset(offset).Find(&content)
	queryBuider.Model(&models.BlogUser{}).Find(&content)

	//mgr.db.Scopes(paginate(categories, &pagination, mgr.db)).Find(&categories)
	return content, len
}

// 查找文章详情
func (mgr *manager) GetEditDetails(id string, userId int) models.Content {
	var content models.Content
	mgr.db.Where("id = ? AND user_id >= ?", id, userId).First(&content)

	//mgr.db.Scopes(paginate(categories, &pagination, mgr.db)).Find(&categories)
	return content
}

// 查找用户资料信息
func (mgr *manager) GetUserInfo(id int) models.UserInfo {
	var content models.UserInfo
	//mgr.db.First(&content)
	mgr.db.Where("user_id=?", id).First(&content)
	return content
}
