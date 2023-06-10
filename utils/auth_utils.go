package utils

import (
	"Jakpat_Test_2/models"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func CheckPasswordMatch(hashPassword string, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(currPassword))

	return err == nil
}

func GenerateToken(request models.User, expired time.Duration) (*string, error) {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	var APPLICATION_NAME = "GOLANG-STORE"

	expiredTime := time.Now().Add(time.Hour * expired).Unix()
	claims := &models.TokenClaims{
		Id:        int16(request.UserID),
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: expiredTime,
		},
	}

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokesStr, err := tokens.SignedString(jwtTokenSecret)
	if err != nil {
		return nil, err
	}
	return &tokesStr, nil

}

func TokenClaims(userToken string) (*models.TokenClaims, error) {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	claims := &models.TokenClaims{}

	token, err := jwt.ParseWithClaims(userToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtTokenSecret, nil
	})
	if err != nil {
		return nil, errors.New("failed claims tokens")
	}

	if !token.Valid {
		return nil, errors.New("tokens not valid")
	}

	return claims, nil
}
