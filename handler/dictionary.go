package handler

import (
	"aceh-dictionary-api/dictionary"
	"aceh-dictionary-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type dictionaryHandler struct {
	service dictionary.Service
}

func NewDictionaryHandler(dictService dictionary.Service) *dictionaryHandler {
	return &dictionaryHandler{dictService}
}

func (h *dictionaryHandler) CreateDictionaryData(c *gin.Context) {
	var input dictionary.AcehIndo

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("error", "Failed to create dictionary data", http.StatusBadRequest, false)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isSuccess, err := h.service.SaveData(input)
	if err != nil {
		response := helper.APIResponse("error", "Failed to create dictionary data", http.StatusBadRequest, false)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !isSuccess {
		response := helper.APIResponse("error", "Failed to create dictionary data", http.StatusBadRequest, false)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success", "Successfully to create dictionary data", http.StatusOK, isSuccess)
	c.JSON(http.StatusOK, response)
}
