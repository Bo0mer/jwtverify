package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func init() {
	log.SetFlags(0)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("usage: jwtverify <token> <key>\n")
	}

	tokenString := os.Args[1]
	keyPath := os.Args[2]

	publicKey, err := readKey(keyPath)
	if err != nil {
		log.Fatalf("error parsing public key: %v\n", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	fmt.Println(token.Valid)
}

func readKey(path string) (*rsa.PublicKey, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	key, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(key)
}
