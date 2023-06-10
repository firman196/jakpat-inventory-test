package usecase

import (
	"Jakpat_Test_2/models"
)

type UserUsecase interface {
	Register(input models.User) (*models.User, error)
	Login(input models.User) (*models.Token, error)
	RefreshToken(refreshToken string) (*models.Token, error)
}
