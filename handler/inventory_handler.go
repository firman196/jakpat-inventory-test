package handler

import "github.com/gin-gonic/gin"

type InventoryHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetById(c *gin.Context)
	GetBySku(c *gin.Context)
	GetBySeller(c *gin.Context)
	DeleteById(c *gin.Context)
}
