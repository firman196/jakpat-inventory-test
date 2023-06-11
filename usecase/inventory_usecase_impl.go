package usecase

import (
	"Jakpat_Test_2/models"
	"Jakpat_Test_2/repository"
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	roleSeller   = "seller"
	roleCustomer = "customer"
)

type InventoryUsecaseImpl struct {
	repository repository.InventoryRepository
}

func NewInventoryUsecaseImpl(repository repository.InventoryRepository) InventoryUsecase {
	return &InventoryUsecaseImpl{
		repository: repository,
	}
}

func (u *InventoryUsecaseImpl) Create(user models.User, input models.InventoryInput) (*models.Inventory, error) {
	if roleSeller != user.Role {
		return nil, errors.New("FORBIDDEN TO ACCESS")
	}
	category := models.Inventory{
		Sku:         uuid.New().String(),
		ProductName: input.ProductName,
		QtyTotal:    input.QtyTotal,
		QtyReserved: input.QtyReserved,
		QtySaleable: input.QtySaleable,
		SellerId:    user.UserID,
	}

	response, err := u.repository.Create(category)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *InventoryUsecaseImpl) Update(user models.User, id int, input models.InventoryInput) (*models.Inventory, error) {
	if roleSeller != user.Role {
		return nil, errors.New("FORBIDDEN TO ACCESS")
	}
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

func (u *InventoryUsecaseImpl) GetById(user models.User, id int) (*models.Inventory, error) {
	if roleSeller != user.Role {
		return nil, errors.New("FORBIDDEN TO ACCESS")
	}
	inventory, err := u.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (u *InventoryUsecaseImpl) GetBySku(user models.User, sku string) (*models.Inventory, error) {
	if roleSeller != user.Role {
		return nil, errors.New("FORBIDDEN TO ACCESS")
	}
	inventory, err := u.repository.FindBySku(sku)
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (u *InventoryUsecaseImpl) GetBySeller(user models.User) ([]models.Inventory, error) {
	if roleSeller != user.Role {
		return nil, errors.New("FORBIDDEN TO ACCESS")
	}
	inventories, err := u.repository.FindBySellerId(int(user.UserID))
	if err != nil {
		return inventories, err
	}

	return inventories, nil
}

func (u *InventoryUsecaseImpl) Delete(user models.User, id int) (bool, error) {
	if roleSeller != user.Role {
		return false, errors.New("FORBIDDEN TO ACCESS")
	}
	inventory, err := u.repository.FindByID(id)
	if err != nil {
		return false, err
	}

	newInventory := models.Inventory{
		Id:          inventory.Id,
		Sku:         inventory.Sku,
		ProductName: inventory.ProductName,
		QtyTotal:    inventory.QtyTotal,
		QtyReserved: inventory.QtyReserved,
		QtySaleable: inventory.QtySaleable,
		SellerId:    inventory.SellerId,
		DeletedBy:   user.UserID,
		DeletedAt:   time.Now(),
		IsDeleted:   true,
	}

	_, errUpdate := u.repository.Update(newInventory)

	if errUpdate != nil {
		return false, err
	}

	return true, nil
}
