package handler

import (
	"Jakpat_Test_2/models"
	"Jakpat_Test_2/usecase"
	"Jakpat_Test_2/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var user = models.User{
	UserID: 1,
	Role:   "seller",
}

type InventoryHandlerImpl struct {
	InventoryUsecase usecase.InventoryUsecase
}

func NewInventoryHandlerImpl(inventoryUsecase usecase.InventoryUsecase) InventoryHandler {
	return &InventoryHandlerImpl{
		InventoryUsecase: inventoryUsecase,
	}
}

// CreateInventory	godoc
// @Summary			Create a new Inventory
// @Description		Save Inventory data in Db.
// @Param			user body models.InventoryInput true "Create inventory"
// @Produce			application/json
// @Tags			inventory
// @Success			200 {object} utils.Response
// @Router			/api/v1/inventory [post]
func (h *InventoryHandlerImpl) Create(c *gin.Context) {
	var input models.InventoryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Create Category Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	category, errService := h.InventoryUsecase.Create(user, input)
	if errService != nil {
		response := utils.ApiResponse("Create category failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Create category success", http.StatusOK, "success", category)
	c.JSON(http.StatusOK, response)
}

// UpdateInventory	godoc
// @Summary			Update Inventory
// @Description		Update data inventory in Db.
// @Param			user body models.InventoryInput true "Update inventory"
// @Produce			application/json
// @Tags			inventory
// @Success			200 {object} utils.Response
// @Router			/api/v1/inventory/{id} [put]
func (h *InventoryHandlerImpl) Update(c *gin.Context) {
	id, errParam := strconv.Atoi(c.Param("id"))

	if errParam != nil {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.InventoryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Update inventory failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	inventory, errService := h.InventoryUsecase.Update(user, id, input)
	if errService != nil {
		response := utils.ApiResponse("Update inventory failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Update inventory success", http.StatusOK, "success", inventory)
	c.JSON(http.StatusOK, response)
}

// FindByIdInventory 	godoc
// @Summary				Get Single inventory by id.
// @Param				id path integer true "find inventory by id"
// @Description			Return data inventory where similar with id.
// @Produce				application/json
// @Tags				inventory
// @Success				200 {object} utils.Response
// @Router				/api/v1/inventory/{id} [get]
func (h *InventoryHandlerImpl) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := utils.ApiResponse("Get data inventory failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inventory, errService := h.InventoryUsecase.GetById(user, id)
	if errService != nil {
		response := utils.ApiResponse("Get data inventory failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data inventory success", http.StatusOK, "success", inventory)
	c.JSON(http.StatusOK, response)
}

// FindBySkuInventory 	godoc
// @Summary				Get Single inventory by sku.
// @Param				sku path string true "fing inventory by sku"
// @Description			Return data inventory where similar with sku.
// @Produce				application/json
// @Tags				inventory
// @Success				200 {object} utils.Response
// @Router				/api/v1/inventory/{id} [get]
func (h *InventoryHandlerImpl) GetBySku(c *gin.Context) {
	sku := c.Param("sku")
	inventory, errService := h.InventoryUsecase.GetBySku(user, sku)
	if errService != nil {
		response := utils.ApiResponse("Get data inventory failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data inventory success", http.StatusOK, "success", inventory)
	c.JSON(http.StatusOK, response)
}

func (h *InventoryHandlerImpl) GetBySeller(c *gin.Context) {
	inventories, errService := h.InventoryUsecase.GetBySeller(user)
	if errService != nil {
		response := utils.ApiResponse("Get all data inventory failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data inventory success", http.StatusOK, "success", inventories)
	c.JSON(http.StatusOK, response)
}

// DeleteByIdInventory 	godoc
// @Summary				Delete inventory by id.
// @Param				id path integer true "delete inventory by id"
// @Description			Return data boolean.
// @Produce				application/json
// @Tags				inventory
// @Success				200 {object} utils.Response
// @Router				/api/v1/inventory/{id} [delete]
func (h *InventoryHandlerImpl) DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := utils.ApiResponse("Get data inventory failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inventory, errService := h.InventoryUsecase.GetById(user, id)
	if errService != nil {
		response := utils.ApiResponse("Get data inventory failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data inventory success", http.StatusOK, "success", inventory)
	c.JSON(http.StatusOK, response)
}
