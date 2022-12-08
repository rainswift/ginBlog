package router

import (
	"fmt"
	"ginBlog/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ClientIp(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" && !IsLocalIp(ip) {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" && !IsLocalIp(ip) {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		if !IsLocalIp(ip) {
			return ip
		}
	}

	return ""

}
func IsLocalIp(ip string) bool {
	/*
		局域网（intranet）的IP地址范围包括：

		10．0．0．0／8－－10．0．0．0～10．255．255．255（A类）

		172．16．0．0／12－172．16．0．0－172．31．255．255（B类）

		192．168．0．0／16－－192．168．0．0～192．168．255．255（C类）
	*/
	ipAddr := strings.Split(ip, ".")

	if strings.EqualFold(ipAddr[0], "10") {
		return true
	} else if strings.EqualFold(ipAddr[0], "172") {
		addr, _ := strconv.Atoi(ipAddr[1])
		if addr >= 16 && addr < 31 {
			return true
		}
	} else if strings.EqualFold(ipAddr[0], "192") && strings.EqualFold(ipAddr[1], "168") {
		return true
	}
	return false
}

func GetRealIp(r *http.Request) string {
	ip := ClientIp(r)
	return ip
}

func MyMiddleware1(c *gin.Context) {
	ip := GetRealIp(c.Request)
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	fmt.Println(ip)
}

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
	e.Use(MyMiddleware1)
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
