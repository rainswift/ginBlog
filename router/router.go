package router

import (
	"ginBlog/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Start() {
	e := gin.Default()
	// 实现跨域访问
	mwCORS := cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	})

	e.Use(mwCORS)
	e.POST("/add", api.AddUser)
	e.POST("/editSave", api.EditSave)
	e.POST("/userDelect", api.UserDelect)
	e.GET("/userList", api.GetUserList)
	e.POST("/login", api.Login)
	e.GET("/edit/list", api.GetEditList)
	e.GET("/edit/deatils", api.GetDeatils)
	e.POST("/edit/delect", api.EditDelect)

	e.POST("/userSave", api.UserSave)
	e.GET("/getUserInfo", api.GetUserInfo)
	e.Run()
}
