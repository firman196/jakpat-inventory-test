package usecase

import "Jakpat_Test_2/models"

type InventoryUsecase interface {
	Create(user models.User, input models.InventoryInput) (*models.Inventory, error)
	Update(user models.User, id int, input models.InventoryInput) (*models.Inventory, error)
	GetById(user models.User, id int) (*models.Inventory, error)
	GetBySku(user models.User, sku string) (*models.Inventory, error)
	GetBySeller(user models.User) ([]models.Inventory, error)
	Delete(user models.User, id int) (bool, error)
}
