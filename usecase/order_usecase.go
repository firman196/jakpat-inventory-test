package usecase

import "Jakpat_Test_2/models"

type OrderUsecase interface {
	Create(user *models.User, input models.OrderInput) (*models.SalesOrder, error)
	Update(user *models.User, id string, input models.OrderInput) (*models.SalesOrder, error)
	GetById(user *models.User, id string) (*models.SalesOrder, error)
	GetBySeller(user *models.User) ([]models.SalesOrder, error)
	SetExpiredOrder()
	Delete(user *models.User, id string) (bool, error)
}
