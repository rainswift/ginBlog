package responose

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ResponseList struct {
	Response
	Total int64 `json:"total"`
}
type ResponseToken struct {
	Response
	Token string `json:"token"`
}

// Success 请求成功返回
func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{200, message, data})
}

func SuccessList(message string, data interface{}, total int64, c *gin.Context) {
	res := Response{200, message, data}
	c.JSON(http.StatusOK, ResponseList{res, total})
}
func SuccessToken(message string, data interface{}, token string, c *gin.Context) {
	res := Response{200, message, data}
	c.JSON(http.StatusOK, ResponseToken{res, token})
}

// Failed 请求失败返回
func Failed(message string, c *gin.Context) {
	c.JSON(400, Response{400, message, 0})
}

func FailedToken(message string, c *gin.Context) {
	c.JSON(403, Response{403, message, 0})
}
