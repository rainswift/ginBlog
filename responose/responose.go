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

// Success 请求成功返回
func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{200, message, data})
}

func SuccessList(message string, data interface{}, total int64, c *gin.Context) {
	res := Response{200, message, data}
	c.JSON(http.StatusOK, ResponseList{res, total})
}

// Failed 请求失败返回
func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{400, message, 0})
}
