package repository

import (
	"Jakpat_Test_2/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) Create(user models.User) (*models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) Update(user models.User) (*models.User, error) {
	err := r.db.Model(&user).Where("email = ?", user.Email).Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Model(&user).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
