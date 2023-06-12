package usecase

import (
	"Jakpat_Test_2/models"
	"Jakpat_Test_2/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var inventoryInput = models.InventoryInput{
	ProductName: "test product",
	QtyTotal:    10,
	QtyReserved: 2,
	QtySaleable: 8,
}

var inventory = models.Inventory{
	Id:          1,
	Sku:         "sku-test",
	ProductName: "test product",
	QtyTotal:    10,
	QtyReserved: 2,
	QtySaleable: 8,
	SellerId:    1,
}

var user = models.User{
	UserID:    1,
	Firstname: "Firman",
	Lastname:  "saputro",
	Email:     "firman@gmail.com",
	Role:      "seller",
}

// Scenario successfully
// testing Create category service using testify and mock
func TestCreateSuccess(t *testing.T) {
	var inventoryRepository = &mocks.InventoryRepositoryMock{Mock: mock.Mock{}}
	var inventoryUsecase = InventoryUsecaseImpl{repository: inventoryRepository}
	//positive case
	inventoryRepository.Mock.On("Create", inventory).Return(inventory, nil)
	category, err := inventoryUsecase.Create(&user, inventoryInput)

	assert.Nil(t, err)
	assert.NotNil(t, category)
	inventoryRepository.Mock.AssertExpectations(t)
}
