package token

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type UserClaims struct {
	UserID int
	jwt.StandardClaims
}

var JWTKEY = []byte("wanghui940318")

func GetToken(userId int) string {
	expireTime := time.Now().Add(6000 * time.Second).Unix()
	claims := &UserClaims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKEY)
	if err != nil {
		log.Fatal(err.Error())
	}
	return tokenString
}
