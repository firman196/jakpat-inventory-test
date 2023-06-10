package models

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	UserID    uint   `gorm:"primaryKey; not null"`
	Firstname string `gorm:"size:100;not null"`
	Lastname  string `gorm:"size:100"`
	Email     string `gorm:"size:100;not null;unique"`
	Password  string `gorm:"size:255,not null"`
	Role      string `gorm:"column:role;type:enum('seller','customer')"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	DeletedBy uint `gorm:"not null; autoincrement:false"`
	IsDeleted bool `gorm:"default:false"`
}

type Inventory struct {
	Id          uint   `gorm:"primaryKey; not null"`
	Sku         string `gorm:"size:100; not null;unique"`
	ProductName string `gorm:"size:255;not null"`
	QtyTotal    int    `gorm:"default:0"`
	QtyReserved int    `gorm:"default:0"`
	QtySaleable int    `gorm:"default:0"`
	SellerId    uint   `gorm:"size:36;not null;uniqueIndex;primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SalesOrder struct {
	OrderId     string `gorm:"primaryKey;size:100; not null"`
	SellerId    uint   `gorm:"not null; autoIncrement:false"`
	CustomerId  uint   `gorm:"not null; autoIncrement:false"`
	InventoryId uint   `gorm:"not null; autoincrement:false"`
	Status      string `gorm:"column:status;type:enum('waiting', 'on_process', 'shipping', 'delivered','expired')"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	DeletedBy   uint `gorm:"not null; autoincrement:false"`
	IsDeleted   bool `gorm:"default:false"`
}
