package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"blog/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string, id string, expires int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(expires) * time.Hour)

	claims := Claims{
		id,
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

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
