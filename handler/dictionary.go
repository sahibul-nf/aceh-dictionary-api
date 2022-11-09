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

func NewDictionaryHandler(dictionaryService dictionary.Service) *dictionaryHandler {
	return &dictionaryHandler{dictionaryService}
}

func (h *dictionaryHandler) AddNewDictionary(c *gin.Context) {
	var input dictionary.DictionaryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Failed to add new dictionary", http.StatusBadRequest, false)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dictionary, err := h.service.SaveData(input)
	if err != nil {
		response := helper.APIResponse("Failed to add new dictionary", http.StatusBadRequest, false)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully to add new dictionary", http.StatusOK, dictionary)
	c.JSON(http.StatusOK, response)
}

func (h *dictionaryHandler) GetDictionaries(c *gin.Context) {

	dictionaries, err := h.service.GetAllData()
	if err != nil {
		response := helper.APIResponse("Failed to get all dictionary", http.StatusBadRequest, dictionaries)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := dictionary.FormatDictionaries(dictionaries)

	response := helper.APIResponse("Successfully to get all dictionary", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}
