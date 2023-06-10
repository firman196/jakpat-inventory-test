package usecase

import (
	"Jakpat_Test_2/models"
	"Jakpat_Test_2/repository"
	"Jakpat_Test_2/utils"
	"errors"
	"os"
	"strconv"
	"time"
)

type UserUsecaseImpl struct {
	repository repository.UserRepository
}

func NewUserUsecaseImpl(repository repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{
		repository: repository,
	}
}

func (u *UserUsecaseImpl) Register(input models.User) (*models.User, error) {
	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	user := models.User{
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Password:  hash,
		Email:     input.Email,
		Role:      input.Role,
	}
	response, err := u.repository.Create(user)
	if err != nil {
		return nil, err
	}

	return response, nil

}

func (u *UserUsecaseImpl) Login(input models.User) (*models.Token, error) {
	user, err := u.repository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	auth := utils.CheckPasswordMatch(user.Password, input.Password)

	if !auth {
		return nil, errors.New("EMAIL & PASSWORD NOT MATCHED")
	}

	jwtExpiredTimeToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_TOKEN"))
	jwtExpiredTimeRefreshToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_REFRESH_TOKEN"))

	tokenCreateRequest := &models.User{
		UserID:    user.UserID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}
	tokens, errToken := utils.GenerateToken(*tokenCreateRequest, time.Duration(jwtExpiredTimeToken))
	refreshTokens, errRefToken := utils.GenerateToken(*tokenCreateRequest, time.Duration(jwtExpiredTimeRefreshToken))
	if errToken != nil || errRefToken != nil {
		return nil, errors.New("GENERATE TOKEN FAILED")
	}
	token := models.Token{
		Token:        *tokens,
		RefreshToken: *refreshTokens,
	}

	return &token, nil
}

func (u *UserUsecaseImpl) RefreshToken(refreshToken string) (*models.Token, error) {
	claims, errClaims := utils.TokenClaims(refreshToken)
	if errClaims != nil {
		return nil, errClaims
	}

	_, err := u.repository.FindByEmail(claims.Email)
	if err != nil {
		return nil, err
	}
	jwtExpiredTimeToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_TOKEN"))
	jwtExpiredTimeRefreshToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_REFRESH_TOKEN"))

	tokenCreateRequest := &models.User{
		UserID:    uint(claims.Id),
		Email:     claims.Email,
		Firstname: claims.Firstname,
		Lastname:  claims.Lastname,
	}

	tokens, errToken := utils.GenerateToken(*tokenCreateRequest, time.Duration(jwtExpiredTimeToken))
	refreshTokens, errRefToken := utils.GenerateToken(*tokenCreateRequest, time.Duration(jwtExpiredTimeRefreshToken))
	if errToken != nil || errRefToken != nil {
		return nil, errors.New("GENERATE TOKEN FAILED")
	}

	token := models.Token{
		Token:        *tokens,
		RefreshToken: *refreshTokens,
	}

	return &token, nil
}
