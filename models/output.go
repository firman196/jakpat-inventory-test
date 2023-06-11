package models

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenClaims struct {
	Id        int16  `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	jwt.StandardClaims
}
