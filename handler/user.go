package handler

import (
	"aceh-dictionary-api/auth"
	"aceh-dictionary-api/helper"
	"aceh-dictionary-api/user"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationFormat(err)

		response := helper.APIResponse("Failed to register user", http.StatusBadRequest, nil, errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Failed to register user", http.StatusConflict, nil, err)
		c.JSON(http.StatusConflict, response)
		return
	}

	formatter := user.FormatUser(newUser)

	response := helper.APIResponse("Successfully to register user", http.StatusOK, formatter, nil)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationFormat(err)

		response := helper.APIResponse("Failed to login", http.StatusBadRequest, nil, errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailNotRegistered := h.userService.IsEmailAvailable(input.Email)
	if isEmailNotRegistered {
		errors := errors.New("email not registered")
		response := helper.APIResponse("Failed to login", http.StatusUnauthorized, nil, errors.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	loggedInUser, err := h.userService.LoginUser(input)
	if err != nil {
		response := helper.APIResponse("Failed to login", http.StatusUnauthorized, nil, err)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		response := helper.APIResponse("Failed to login", http.StatusInternalServerError, nil, err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := user.FormatAuthUser(token)

	response := helper.APIResponse("Successfully to login", http.StatusOK, formatter, nil)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUserInfo(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	authUser, err := h.userService.GetUserByID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user", http.StatusInternalServerError, nil, err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := user.FormatUser(authUser)

	response := helper.APIResponse("Successfully to get user", http.StatusOK, formatter, nil)
	c.JSON(http.StatusOK, response)
}
