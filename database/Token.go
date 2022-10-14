package database

import (
	"ShopProject/errs"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

const TokenKey = "hellotheregeneralkenobi"

type Claims struct {
	jwt.StandardClaims
	IsAdmin string `json:"is_admin"`
}

func GenerateTokenForAdmin() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		IsAdmin: "True",
	})

	return token.SignedString([]byte(TokenKey))
}

func ParseToken(tokenString string) (error, string) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("IsAdmin"), nil
	})
	if err != nil {
		errs.Printer(err, "ParseToken1")
	}

	myClaims := token.Claims.(*Claims)

	return err, myClaims.IsAdmin
}

func GenerateTokenForUser() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		IsAdmin: "False",
	})

	return token.SignedString([]byte(TokenKey))
}
