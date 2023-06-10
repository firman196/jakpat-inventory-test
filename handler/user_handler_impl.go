package handler

import (
	"Jakpat_Test_2/models"
	"Jakpat_Test_2/usecase"
	"Jakpat_Test_2/utils"
	"golang-store/model/web"
	"net/http"
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
// @Param			user body web.CreateUser true "Create user"
// @Produce			application/json
// @Tags			user
// @Success			200 {object} utils.Response
// @Router			/api/v1/user/register [post]
func (service *UserHandlerImpl) Register(c *gin.Context) {
	var input models.User
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Register Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	user, errService := service.UserService.Register(input)
	if errService != nil {
		response := utils.ApiResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Register user success", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}

func (service *UserControllerImpl) Login(c *gin.Context) {
	var input web.Login

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, errService := service.UserService.Login(input)
	if errService != nil {
		response := utils.ApiResponse("Login user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Login user success", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}

func (service *UserControllerImpl) RefreshToken(c *gin.Context) {

}
