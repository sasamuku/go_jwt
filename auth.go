package main

import (
	"log"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
)

var hmacSecret = []byte(os.Getenv("SIGNINGKEY"))

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "1234567890"
	claims["name"] = "Mike"
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.Fatal("tokenString:", err)
	}

	w.Write([]byte(tokenString))
})

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
