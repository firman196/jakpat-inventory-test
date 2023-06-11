package usecase

import (
	"Jakpat_Test_2/models"
	"Jakpat_Test_2/repository"

	"github.com/google/uuid"
)

type InventoryUsecaseImpl struct {
	repository repository.InventoryRepository
}

func NewInventoryUsecaseImpl(repository repository.InventoryRepository) InventoryUsecase {
	return &InventoryUsecaseImpl{
		repository: repository,
	}
}

func (u *InventoryUsecaseImpl) Create(input models.InventoryInput) (*models.Inventory, error) {
	category := models.Inventory{
		Sku:         uuid.New().String(),
		ProductName: input.ProductName,
		QtyTotal:    input.QtyTotal,
		QtyReserved: input.QtyReserved,
		QtySaleable: input.QtySaleable,
		SellerId:    0,
	}

	response, err := u.repository.Create(category)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *InventoryUsecaseImpl) Update(id int, input models.InventoryInput) (*models.Inventory, error) {
	inventory, err := u.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	newInventory := models.Inventory{
		Id:          inventory.Id,
		Sku:         inventory.Sku,
		ProductName: input.ProductName,
		QtyTotal:    input.QtyTotal,
		QtyReserved: input.QtyReserved,
		QtySaleable: input.QtySaleable,
		SellerId:    inventory.SellerId,
	}

	response, err := u.repository.Update(newInventory)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *InventoryUsecaseImpl) GetById(id int) (*models.Inventory, error) {
	inventory, err := u.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (u *InventoryUsecaseImpl) GetBySku(sku string) (*models.Inventory, error) {
	inventory, err := u.repository.FindBySku(sku)
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (u *InventoryUsecaseImpl) GetAll() ([]models.Inventory, error) {
	inventories, err := u.repository.FindAll()
	if err != nil {
		return inventories, err
	}

	return inventories, nil
}
