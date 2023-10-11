package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var Secret = []byte("xyz")

// jwt过期时间
const expiration = time.Hour * 2

type Claims struct {
	UserID int
	jwt.StandardClaims
}

func GenToken(userid int) (string, error) {
	//创建声明
	a := Claims{
		UserID: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gin-jwt-demo",
			Id:        "",
			NotBefore: 0,
			Subject:   "",
		},
	}

	//哈希方法创建签名
	tt:=jwt.NewWithClaims(jwt.SigningMethodHS256,a)
	tokenString, err := tt.SignedString(Secret)
    if err != nil {
        return "", err
    }
    return tokenString, nil

}

func ParseToken(tokenStr string) (*Claims, error) {
	if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
		tokenStr = tokenStr[7:]
	}
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	//检验token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
