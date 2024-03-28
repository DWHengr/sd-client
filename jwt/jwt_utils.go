package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	tokenSecret = "gk-sd-cloud"
)

// GenerateJWTToken 生成 JWT 令牌
func GenerateJWTToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(10 * time.Minute).Unix(),
	})

	return token.SignedString([]byte(tokenSecret))
}

// VerifyJWTToken 验证 JWT 令牌
func VerifyJWTToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证令牌的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// 返回用于验证令牌的密钥
		return []byte(tokenSecret), nil
	})
}
