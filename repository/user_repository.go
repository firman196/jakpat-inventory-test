package repository

import "Jakpat_Test_2/models"

type UserRepository interface {
	Create(user models.User) (*models.User, error)
	Update(user models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}
