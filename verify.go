package main

import (
	"errors"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

func verifyToken(token, keyFile, secret string) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		switch token.Method.(type) {
		case *jwt.SigningMethodRSA:
			keyBytes, err := ioutil.ReadFile(keyFile)
			if err != nil {
				return nil, err
			}
			return jwt.ParseRSAPublicKeyFromPEM(keyBytes)
		case *jwt.SigningMethodHMAC:
			return []byte(secret), nil
		default:
			return nil, errors.New("unsupported signing method")
		}
	})

	if err != nil {
		if verr, ok := err.(*jwt.ValidationError); ok {
			if verr.Errors&jwt.ValidationErrorMalformed != 0 {
				fail("token is malformed")
			}
			if verr.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				fail("token signature is invalid")
			}
			if verr.Errors&jwt.ValidationErrorExpired != 0 {
				fail("token is expired")
			}
		}
	}

	if !t.Valid {
		fail("token is invalid")
	}
}
