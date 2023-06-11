package repository

import (
	"Jakpat_Test_2/models"
)

type InventoryRepository interface {
	Create(category models.Inventory) (*models.Inventory, error)
	Update(category models.Inventory) (*models.Inventory, error)
	FindByID(id int) (*models.Inventory, error)
	FindBySku(sku string) (*models.Inventory, error)
	FindBySellerId(sellerId int) ([]models.Inventory, error)
}
