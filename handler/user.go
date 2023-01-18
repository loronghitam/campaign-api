package handler

import (
	"campaign-api/helper"
	"campaign-api/user/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService usecase.Service
}

func NewUserHandler(userService usecase.Service) *userHandler {
	return &userHandler{userService}
}

func (h userHandler) RegisterUser(c *gin.Context) {
	var input usecase.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ApiResponseError(err)
		errorMissage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register Account failed", http.StatusUnprocessableEntity, "Error", errorMissage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	user, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Register Account faild", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	format := usecase.NewFormatUser(user, "tokentokentoken")
	response := helper.ApiResponse("Account has been registered", http.StatusCreated, "Success", format)
	c.JSON(http.StatusOK, response)
}
