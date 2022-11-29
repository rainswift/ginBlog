package api

import (
	"ginBlog/models"
	jwt "github.com/dgrijalva/jwt-go"
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
