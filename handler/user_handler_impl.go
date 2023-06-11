package handler

import (
	"Jakpat_Test_2/models"
	"Jakpat_Test_2/usecase"
	"Jakpat_Test_2/utils"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlerImpl struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandlerImpl(userUsecase usecase.UserUsecase) UserHandler {
	return &UserHandlerImpl{
		UserUsecase: userUsecase,
	}
}

// RegisterUser		godoc
// @Summary			Register a user
// @Description		Save user data in Db.
// @Param			user body models.RegisterInput true "Create user"
// @Produce			application/json
// @Tags			user
// @Success			200 {object} utils.Response
// @Router			/api/v1/user/register [post]
func (h *UserHandlerImpl) Register(c *gin.Context) {
	var input models.RegisterInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Register Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	user, errService := h.UserUsecase.Register(input)
	if errService != nil {
		response := utils.ApiResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Register user success", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}

// LoginUser		godoc
// @Summary			Login a user
// @Description		Authenticate User.
// @Param			user body models.LoginInput true "Login user"
// @Produce			application/json
// @Tags			user
// @Success			200 {object} utils.Response
// @Router			/api/v1/user/login [post]
func (h *UserHandlerImpl) Login(c *gin.Context) {
	var input models.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, errService := h.UserUsecase.Login(input)
	if errService != nil {
		response := utils.ApiResponse("Login user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Login user success", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandlerImpl) RefreshToken(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		response := utils.ApiResponse("User Unauthorized", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token := strings.Split(header, " ")
	if len(token) != 2 {
		response := utils.ApiResponse("User Unauthorized", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if token[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	newToken, errToken := h.UserUsecase.RefreshToken(token[1])

	if errToken != nil {
		response := utils.ApiResponse("User Unauthorized", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Refresh token success", http.StatusOK, "success", newToken)
	c.JSON(http.StatusOK, response)

}
