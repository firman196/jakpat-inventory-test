package usecase

import "Jakpat_Test_2/models"

type InventoryUsecase interface {
	Create(input models.InventoryInput) (*models.Inventory, error)
	Update(id int, input models.InventoryInput) (*models.Inventory, error)
	GetById(id int) (*models.Inventory, error)
	GetBySku(sku string) (*models.Inventory, error)
	GetAll() ([]models.Inventory, error)
}
