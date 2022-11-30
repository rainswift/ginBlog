package api

import (
	"errors"
	"ginBlog/models"
	response "ginBlog/responose"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func GenToken(username string, password string) (string, error) {
	// 创建一个我们自己的声明
	c := models.BlogUser{
		username, // 自定义字段
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "laoguo",                             // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	// 注意这个地方一定要是字节切片不能是字符串
	return token.SignedString([]byte("secret"))
}

var MySecret = []byte("secret")

func ParseToken(c *gin.Context) (*models.BlogUser, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		response.Failed("token过期", c)
		return nil, errors.New("invalid token1")
	}
	tokenString = tokenString[7:]
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &models.BlogUser{},
		func(token *jwt.Token) (i interface{}, err error) {
			return MySecret, nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*models.BlogUser); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
