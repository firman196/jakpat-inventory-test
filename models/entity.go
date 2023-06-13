package models

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	UserID    uint      `gorm:"primaryKey; not null" json:"user_id"`
	Firstname string    `gorm:"size:100;not null" json:"firstname" binding:"required"`
	Lastname  string    `gorm:"size:100" json:"lastname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email" binding:"required,email"`
	Password  string    `gorm:"size:255,not null" json:"password" binding:"required"`
	Role      string    `gorm:"column:role;type:enum('seller','customer')" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"default:null" json:"deleted_at"`
	DeletedBy uint      `gorm:"default:null; autoincrement:false" json:"deleted_by"`
	IsDeleted bool      `gorm:"default:false" json:"is_deleted"`
}

type Inventory struct {
	Id          uint      `gorm:"primaryKey; not null" json:"id"`
	Sku         string    `gorm:"size:100; not null;unique" json:"sku"`
	ProductName string    `gorm:"size:255;not null" json:"product_name"`
	QtyTotal    int       `gorm:"default:0" json:"qty_total"`
	QtyReserved int       `gorm:"default:0" json:"qty_reserved"`
	QtySaleable int       `gorm:"default:0" json:"qty_saleable"`
	SellerId    uint      `gorm:"size:36;not null" json:"seller_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `gorm:"default:null" json:"deleted_at"`
	DeletedBy   uint      `gorm:"default:null; autoincrement:false" json:"deleted_by"`
	IsDeleted   bool      `gorm:"default:false" json:"is_deleted"`
}

type SalesOrder struct {
	OrderId         string    `gorm:"primaryKey;size:100; not null" json:"order_id"`
	CustomerId      uint      `gorm:"not null; autoIncrement:false" json:"customer_id"`
	InventoryId     uint      `gorm:"not null; autoincrement:false" json:"inventory_id"`
	QtyOrder        int       `gorm:"default:0" json:"qty_order"`
	ShippingAddress string    `gorm:"not null; autoincrement:false" json:"shipping_address"`
	NoTelphone      string    `gorm:"not null; autoincrement:false" json:"no_telphone"`
	Status          string    `gorm:"column:status;type:enum('waiting', 'on_process', 'shipping', 'delivered','expired')" json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ExpiredAt       time.Time `json:"expired_at"`
}
