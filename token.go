package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var key = []byte("THIS_IS_KEY")

func newJWT() (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)

	claims := jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func main() {
	t, err := newJWT()
	if err != nil {
		println(err.Error())
	} else {
		println(t)
	}
}
