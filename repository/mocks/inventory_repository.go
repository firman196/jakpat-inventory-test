package mocks

import (
	"Jakpat_Test_2/models"

	mock "github.com/stretchr/testify/mock"
)

type InventoryRepositoryMock struct {
	mock.Mock
}

func (r *InventoryRepositoryMock) Create(inventory models.Inventory) (*models.Inventory, error) {
	args := r.Called(inventory)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		newInventory := args.Get(0).(models.Inventory)
		return &newInventory, nil
	}
}

func (r *InventoryRepositoryMock) Update(inventory models.Inventory) (*models.Inventory, error) {
	args := r.Called(inventory)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(models.Inventory)
		return &result, nil
	}
}

func (r *InventoryRepositoryMock) FindByID(id int) (*models.Inventory, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(models.Inventory)
		return &result, nil
	}
}

func (r *InventoryRepositoryMock) FindBySku(sku string) (*models.Inventory, error) {
	args := r.Called(sku)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(models.Inventory)
		return &result, nil
	}
}

func (r *InventoryRepositoryMock) FindBySellerId(sellerId int) ([]models.Inventory, error) {
	args := r.Called(sellerId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).([]models.Inventory)
		return result, nil
	}
}
