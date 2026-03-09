package jwtCus

import (
	"KimJin/src/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

// 签名密钥
var secret = []byte(config.GlobalConfig.Jwt.Key)

// GenToken 生成JWT
func GenToken(userID uint, username string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
