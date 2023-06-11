package repository

import (
	"Jakpat_Test_2/models"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepositoryImpl(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		db: db,
	}
}

func (r *OrderRepositoryImpl) Create(order models.SalesOrder) (*models.SalesOrder, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepositoryImpl) Update(order models.SalesOrder) (*models.SalesOrder, error) {
	err := r.db.Model(&order).Where("order_id = ?", order.OrderId).Where("is_deleted =?", false).Updates(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepositoryImpl) FindByID(id string) (*models.SalesOrder, error) {
	var order models.SalesOrder
	err := r.db.Model(&order).Where("order_id =?", id).First(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepositoryImpl) FindBySellerId(sellerId int) ([]models.SalesOrder, error) {
	var order []models.SalesOrder

	err := r.db.Where("seller_id =?", sellerId).Preload("User").Find(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
