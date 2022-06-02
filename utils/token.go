package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 24 //设置过期时间


func GenToken(account string) (string, error) {
	// Create the Claims
	 claims := MyClaims{
		 account,
	 	jwt.StandardClaims{
	 		ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
	 		Issuer:    "tony",
		},
	 }
	 token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	 return  token.SignedString([]byte(Config.Secret))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(Config.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
