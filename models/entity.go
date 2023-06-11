package models

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	UserID    uint   `gorm:"primaryKey; not null" json:"user_id"`
	Firstname string `gorm:"size:100;not null" json:"firstname" binding:"required"`
	Lastname  string `gorm:"size:100" json:"lastname"`
	Email     string `gorm:"size:100;not null;unique" json:"email" binding:"required,email"`
	Password  string `gorm:"size:255,not null" json:"password" binding:"required"`
	Role      string `gorm:"column:role;type:enum('seller','customer')" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"default:null"`
	DeletedBy uint      `gorm:"default:null; autoincrement:false"`
	IsDeleted bool      `gorm:"default:false"`
}

type Inventory struct {
	Id          uint   `gorm:"primaryKey; not null"`
	Sku         string `gorm:"size:100; not null;unique"`
	ProductName string `gorm:"size:255;not null"`
	QtyTotal    int    `gorm:"default:0"`
	QtyReserved int    `gorm:"default:0"`
	QtySaleable int    `gorm:"default:0"`
	SellerId    uint   `gorm:"size:36;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"default:null"`
	DeletedBy   uint      `gorm:"default:null; autoincrement:false"`
	IsDeleted   bool      `gorm:"default:false"`
}

type SalesOrder struct {
	OrderId         string `gorm:"primaryKey;size:100; not null"`
	CustomerId      uint   `gorm:"not null; autoIncrement:false"`
	InventoryId     uint   `gorm:"not null; autoincrement:false"`
	ShippingAddress string `gorm:"not null; autoincrement:false"`
	NoTelphone      string `gorm:"not null; autoincrement:false"`
	Status          string `gorm:"column:status;type:enum('waiting', 'on_process', 'shipping', 'delivered','expired')"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
