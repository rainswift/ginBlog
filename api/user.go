package api

import (
	"fmt"
	"ginBlog/dao"
	"ginBlog/models"
	response "ginBlog/responose"
	"ginBlog/utils"
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

func Login(c *gin.Context) {
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
	token, _ := GenToken(user.Username, user.Password)
	fmt.Printf(token)
	if ComparePwd(loadUser.Password, user.Password) {
		fmt.Println("登录成功")
		response.SuccessToken("登录成功", user, token, c)
	} else {
		fmt.Println("密码错误")
		response.Failed("密码错误", c)
	}
}

func EditSave(c *gin.Context) {
	if _, err := ParseToken(c); err != nil {
		return
	}

	var context models.Content
	if err := c.ShouldBind(&context); err != nil {
		response.Failed("参数错误", c)
		return
	}
	dao.Mgr.SaveEdit(&context)
	response.Success("保存成功", context, c)
}

func GetEditList(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	if _, err := ParseToken(c); err != nil {
		fmt.Println(err)
		return
	}
	users, len := dao.Mgr.GetEditList(&pagination)
	response.SuccessList("查询成功5", users, len, c)
}

func GetUserList(c *gin.Context) {
	if _, err := ParseToken(c); err != nil {
		return
	}
	pagination := utils.GeneratePaginationFromRequest(c)
	users := dao.Mgr.GetUserList(&pagination)
	response.Success("查询成功", users, c)
}

func UserDelect(c *gin.Context) {
	//没有完成
	if _, err := ParseToken(c); err != nil {
		return
	}

	pagination := utils.RequestId(c)
	id := c.DefaultPostForm("id", "0")
	//id := c.PostForm("id")

	fmt.Printf("删除id： % + v\n", id)
	users := dao.Mgr.UserDelete(pagination)
	response.Success("删除成功", users, c)
}

func GetDeatils(c *gin.Context) {
	if _, err := ParseToken(c); err != nil {
		return
	}

	var cid models.GetId
	if err := c.ShouldBind(&cid); err != nil {
		response.Failed("参数错误", c)
		return
	}
	users := dao.Mgr.GetEditDetails(cid.Id)
	response.Success("查询成功", users, c)
}

// 个人信息保存
func UserSave(c *gin.Context) {
	if _, err := ParseToken(c); err != nil {
		return
	}

	var context models.UserInfo
	if err := c.ShouldBind(&context); err != nil {
		response.Failed("参数错误", c)
		return
	}

	dao.Mgr.UserSave(&context)
	response.Success("保存成功", context, c)
}

func GetUserInfo(c *gin.Context) {
	if _, err := ParseToken(c); err != nil {
		return
	}
	users := dao.Mgr.GetUserInfo()
	response.Success("查询成功", users, c)
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
