package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	// var test_a int = setting.AppSetting.JwtExpire
	fmt.Println(setting.AppSetting)
	var JwtExpire int = setting.AppSetting.JwtExpire
	expireTime := nowTime.Add(time.Duration(JwtExpire) * time.Hour)

	claims := Claims{
		EncodeMD5(username),
		EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
