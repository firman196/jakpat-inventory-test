package repository

import (
	"Jakpat_Test_2/models"
)

type InventoryRepository interface {
	Create(category models.Inventory) (*models.Inventory, error)
	Update(category models.Inventory) (*models.Inventory, error)
	FindByID(id int) (*models.Inventory, error)
	FindAll() ([]models.Inventory, error)
}
