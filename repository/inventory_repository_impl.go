package repository

import (
	"Jakpat_Test_2/models"

	"gorm.io/gorm"
)

type InventoryRepositoryImpl struct {
	db *gorm.DB
}

func NewInventoryRepositoryImpl(db *gorm.DB) InventoryRepository {
	return &InventoryRepositoryImpl{
		db: db,
	}
}

func (r *InventoryRepositoryImpl) Create(inventory models.Inventory) (*models.Inventory, error) {
	err := r.db.Create(&inventory).Error
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *InventoryRepositoryImpl) Update(inventory models.Inventory) (*models.Inventory, error) {
	err := r.db.Model(&inventory).Where("id = ?", inventory.Id).Updates(&inventory).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (r *InventoryRepositoryImpl) FindByID(id int) (*models.Inventory, error) {
	var inventory models.Inventory
	err := r.db.Model(&inventory).Where("id =?", id).First(&inventory).Error
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *InventoryRepositoryImpl) FindBySku(id string) (*models.Inventory, error) {
	var inventory models.Inventory
	err := r.db.Model(&inventory).Where("sku =?", id).First(&inventory).Error
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *InventoryRepositoryImpl) FindBySellerId(sellerId int) ([]models.Inventory, error) {
	var inventories []models.Inventory

	err := r.db.Where("seller_id =?", sellerId).Preload("User").Find(&inventories).Error
	if err != nil {
		return inventories, err
	}

	return inventories, nil
}
