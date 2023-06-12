package usecase

import (
	"Jakpat_Test_2/models"
	"Jakpat_Test_2/repository"
	"errors"

	"github.com/google/uuid"
)

type OrderUsecaseImpl struct {
	orderRepository     repository.OrderRepository
	inventoryRepository repository.InventoryRepository
}

func NewOrderUsecaseImpl(orderRepository repository.OrderRepository, inventoryRepository repository.InventoryRepository) OrderUsecase {
	return &OrderUsecaseImpl{
		orderRepository:     orderRepository,
		inventoryRepository: inventoryRepository,
	}
}

func (u *OrderUsecaseImpl) Create(user *models.User, input models.OrderInput) (*models.SalesOrder, error) {
	if user.Role != roleCustomer {
		return nil, errors.New("forbidden to access")
	}

	inventory, err := u.inventoryRepository.FindBySku(input.Sku)
	if err != nil {
		return nil, errors.New("sku not found")
	}

	order := models.SalesOrder{
		OrderId:         uuid.New().String(),
		CustomerId:      user.UserID,
		InventoryId:     inventory.Id,
		ShippingAddress: input.ShippingAddress,
		NoTelphone:      input.NoTelphone,
		Status:          input.Status,
	}

	response, err := u.orderRepository.Create(order)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *OrderUsecaseImpl) Update(user *models.User, id string, input models.OrderInput) (*models.SalesOrder, error) {
	if user.Role != roleCustomer {
		return nil, errors.New("forbidden to access")
	}
	orderOld, err := u.orderRepository.FindByID(id)
	if err != nil {
		return nil, errors.New("sku not found")
	}

	order := models.SalesOrder{
		OrderId:         orderOld.OrderId,
		CustomerId:      orderOld.CustomerId,
		InventoryId:     orderOld.InventoryId,
		ShippingAddress: orderOld.ShippingAddress,
		NoTelphone:      orderOld.NoTelphone,
		Status:          input.Status,
	}

	response, err := u.orderRepository.Update(order)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *OrderUsecaseImpl) GetById(user *models.User, id string) (*models.SalesOrder, error) {
	if roleSeller != user.Role {
		return nil, errors.New("forbidden to access")
	}
	order, err := u.orderRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (u *OrderUsecaseImpl) GetBySeller(user *models.User) ([]models.SalesOrder, error) {
	if roleSeller != user.Role {
		return nil, errors.New("forbidden to access")
	}
	orders, err := u.orderRepository.FindBySellerId(int(user.UserID))
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (u *OrderUsecaseImpl) Delete(user *models.User, id string) (bool, error) {
	if roleSeller != user.Role {
		return false, errors.New("forbidden to access")
	}

	_, err := u.orderRepository.DeleteById(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
