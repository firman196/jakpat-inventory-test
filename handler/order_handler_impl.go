package handler

import (
	"Jakpat_Test_2/models"
	"Jakpat_Test_2/usecase"
	"Jakpat_Test_2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandlerImpl struct {
	OrderUsecase usecase.OrderUsecase
}

func NewOrderHandlerImpl(orderUsecase usecase.OrderUsecase) OrderHandler {
	return &OrderHandlerImpl{
		OrderUsecase: orderUsecase,
	}
}

// CreateOrder		godoc
// @Summary			Create a new Order
// @Description		Create new order by customer.
// @Param			order body models.OrderInput true "Create order"
// @Produce			application/json
// @Tags			order
// @Success			200 {object} utils.Response
// @Router			/api/v1/order [post]
func (h *OrderHandlerImpl) Create(c *gin.Context) {
	var input models.OrderInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Order Category Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	order, errService := h.OrderUsecase.Create(user, input)
	if errService != nil {
		response := utils.ApiResponse("Create order failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Create order success", http.StatusOK, "success", order)
	c.JSON(http.StatusOK, response)
}

// UpdateOrder		godoc
// @Summary			Update Order
// @Description		Update status order.
// @Param			user body models.OrderInput true "Update status order"
// @Param			id path integer true "find order by id"
// @Produce			application/json
// @Tags			order
// @Success			200 {object} utils.Response
// @Router			/api/v1/order/{id} [put]
func (h *OrderHandlerImpl) Update(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.OrderInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Update order failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	order, errService := h.OrderUsecase.Update(user, id, input)
	if errService != nil {
		response := utils.ApiResponse("Update order failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Update order success", http.StatusOK, "success", order)
	c.JSON(http.StatusOK, response)
}

// FindByIdOrder 		godoc
// @Summary				Get Single order by id.
// @Param				id path string true "find order by id"
// @Description			Return data order where similar with id.
// @Produce				application/json
// @Tags				order
// @Success				200 {object} utils.Response
// @Router				/api/v1/order/{id} [get]
func (h *OrderHandlerImpl) GetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, errService := h.OrderUsecase.GetById(user, id)
	if errService != nil {
		response := utils.ApiResponse("Get data order failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data order success", http.StatusOK, "success", order)
	c.JSON(http.StatusOK, response)
}

// FindOrderBySeller 	godoc
// @Summary				Get all order by seller.
// @Description			Return data order by seller.
// @Produce				application/json
// @Tags				order
// @Success				200 {object} utils.Response
// @Router				/api/v1/order [get]
func (h *OrderHandlerImpl) GetBySeller(c *gin.Context) {
	orders, errService := h.OrderUsecase.GetBySeller(user)
	if errService != nil {
		response := utils.ApiResponse("Get all data order failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data order success", http.StatusOK, "success", orders)
	c.JSON(http.StatusOK, response)
}

// DeleteByIdOrder 	godoc
// @Summary				Delete order by id.
// @Param				id path string true "delete order by id"
// @Description			Return data boolean.
// @Produce				application/json
// @Tags				order
// @Success				200 {object} utils.Response
// @Router				/api/v1/order/delete/{id} [delete]
func (h *OrderHandlerImpl) DeleteById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inventory, errService := h.OrderUsecase.GetById(user, id)
	if errService != nil {
		response := utils.ApiResponse("Delete order failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Delete order success", http.StatusOK, "success", inventory)
	c.JSON(http.StatusOK, response)
}
