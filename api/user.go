package api

import (
	"fmt"
	"ginBlog/dao"
	"ginBlog/models"
	response "ginBlog/responose"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	r dao.Manager
}

func AddUser(c *gin.Context) {

	var user models.BlogUser

	if err := c.ShouldBind(&user); err != nil {
		response.Failed("参数错误", c)
		return
	}
	if err := AddUser2(user.Username); err != nil {
		response.Failed("用户名已经存在", c)
		return
	}
	b, _ := GenPwd(user.Password)
	user.Password = string(b)
	fmt.Printf("输入的名字 % + v\n", user.Username)
	fmt.Println(c)
	dao.Mgr.AddUser(&user)
	response.Success("添加成功", user, c)
}

// 创建用户
func AddUser2(name string) error {

	// 用户名存在
	flag := dao.Mgr.GetByName(name)
	fmt.Println(flag)
	if flag {
		return nil
	} else {
		return ErrUserExistWithName
	}
	// 无效用户名
	//if models.ValidateUserName(user.Username) {
	//	return ErrInvalidUsername
	//}
	// 无效密码
	//if models.ValidatePassword(user.Password) {
	//	return ErrInvalidPassword
	//}
	// 创建用户
	//c.AddUser(&user)
}

func Load(c *gin.Context) {
	var user models.BlogUser
	if err := c.ShouldBind(&user); err != nil {
		response.Failed("参数错误", c)
		return
	}
	loadUser, isUser := dao.Mgr.GetLoadUser(user.Username)
	if isUser {
		response.Failed("用户名不存在", c)
		return
	}
	if ComparePwd(loadUser.Password, user.Password) {
		fmt.Println("登录成功")
		response.Success("登录成功", user, c)
	} else {
		fmt.Println("密码错误")
		response.Failed("密码错误", c)
	}
}

func EditSave(c *gin.Context) {
	fmt.Println(c)
}

func GetUserList(c *gin.Context) {
	fmt.Printf("列表数据： % + v\n", c)
	pagination := utils.GeneratePaginationFromRequest(c)
	users := dao.Mgr.GetUserList(&pagination)
	response.Success("查询成功", users, c)
}

func UserDelect(c *gin.Context) {
	pagination := utils.RequestId(c)
	//bodyByts, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("删除数据： % + v\n", pagination)
	id := c.DefaultPostForm("id", "0")
	//id := c.PostForm("id")

	fmt.Printf("删除id： % + v\n", id)
	users := dao.Mgr.UserDelete(pagination)
	response.Success("删除成功", users, c)
}

//登录
func Login(c *gin.Context) {

	var user models.BlogUser

	if err := c.ShouldBind(&user); err != nil {
		response.Failed("参数错误", c)
		return
	}

	if err := AddUser2(user.Username); err != nil {
		response.Failed("用户名已经存在", c)
		return
	}
	b, _ := GenPwd(user.Password)
	user.Password = string(b)

	dao.Mgr.AddUser(&user)
	response.Success("添加成功", user, c)
}

// 生成密码
func GenPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost) //加密处理
	return hash, err
}

// 比对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}