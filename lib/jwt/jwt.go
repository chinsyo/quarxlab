package jwt

import (
	"time"
	xjwt "github.com/dgrijalva/jwt-go"	
)

var jwtSecret = []byte("quarxlab_secret")

type Claims struct {
	xjwt.StandardClaims
	UserID uint 
}

func GenerateToken(userID uint) (string, error) {
	createTime := time.Now()
	expireTime := createTime.Add(3 * time.Hour)

	claims := Claims{
		xjwt.StandardClaims{
			ExpiresAt : expireTime.Unix(),
			Issuer : "quarxlab",
		},
		userID,
	}

	tokenClaims := xjwt.NewWithClaims(xjwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := xjwt.ParseWithClaims(token, &Claims{}, func(token *xjwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}