package usecase

import (
	"Jakpat_Test_2/models"
)

type UserUsecase interface {
	Register(input models.RegisterInput) (*models.User, error)
	Login(input models.LoginInput) (*models.Token, error)
	RefreshToken(refreshToken string) (*models.Token, error)
}
