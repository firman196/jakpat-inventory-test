package models

type RegisterInput struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type InventoryInput struct {
	ProductName string `json:"product_name" binding:"required"`
	QtyTotal    int    `json:"qty_total"`
	QtyReserved int    `json:"qty_reserved"`
	QtySaleable int    `json:"qty_saleable"`
}

type OrderInput struct {
	Sku             string `json:"sku" binding:"required"`
	QtyOrder        int    `json:"qty_order" binding:"required"`
	ShippingAddress string `json:"shipping_address" binding:"required"`
	NoTelphone      string `json:"no_telphone" binding:"required"`
	Status          string `json:"status" binding:"required"`
}
