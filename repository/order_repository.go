package repository

import "Jakpat_Test_2/models"

type OrderRepository interface {
	Create(order models.SalesOrder) (*models.SalesOrder, error)
	Update(order models.SalesOrder) (*models.SalesOrder, error)
	FindByID(id string) (*models.SalesOrder, error)
	FindBySellerId(sellerId int) ([]models.SalesOrder, error)
}
