package handler

import (
	"bwa-golang/helper"
	"bwa-golang/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMsg := gin.H{"errors": errors}

		respone := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, respone)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMsg := gin.H{"errors": errors}
		respone := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", errorMsg)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	formatter := user.FormatUser(newUser, "JWT")

	respone := helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, respone)
}

func (h* userHandler) Login(c *gin.Context){
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMsg := gin.H{"errors": errors}

		respone := helper.APIResponse("Login Account Failed", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, respone)
		return
	}

	logginedUser , err := h.userService.LoginUser(input)

	if err != nil {
		errorMsg := gin.H{"errors": err.Error()}

		respone := helper.APIResponse("Login Account Failed", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, respone)
		return
	}

	formatter := user.FormatUser(logginedUser, "JWT")
	respone := helper.APIResponse("Account has been Logined", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, respone)
}