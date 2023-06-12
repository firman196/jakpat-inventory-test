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
