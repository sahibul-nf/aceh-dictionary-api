package handler

import (
	"aceh-dictionary-api/dictionary"
	"aceh-dictionary-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type dictionaryHandler struct {
	service dictionary.Service
}

func NewDictionaryHandler(dictionaryService dictionary.Service) *dictionaryHandler {
	return &dictionaryHandler{dictionaryService}
}

func (h *dictionaryHandler) AddNewWord(c *gin.Context) {
	var input dictionary.DictionaryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationFormat(err)

		response := helper.APIResponse("Failed to add new word", http.StatusBadRequest, nil, errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dictionary, err := h.service.SaveWord(input)
	if err != nil {
		errors := err

		response := helper.APIResponse("Failed to add new word", http.StatusInternalServerError, dictionary, errors)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Successfully to add new word", http.StatusOK, dictionary, nil)
	c.JSON(http.StatusOK, response)
}

func (h *dictionaryHandler) GetWords(c *gin.Context) {

	dictionaries, err := h.service.GetWords()
	if err != nil {
		errors := err

		response := helper.APIResponse("Failed to get all word", http.StatusInternalServerError, dictionaries, errors)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := dictionary.FormatDictionariesWithTotal(dictionaries)

	response := helper.APIResponse("Successfully to get all word", http.StatusOK, data, nil)
	c.JSON(http.StatusOK, response)
}

func (h *dictionaryHandler) GetWord(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	if id < 1 {
		errors := "id param is invalid, and must be greater than 0"

		response := helper.APIResponse("Failed to get word detail", http.StatusBadRequest, nil, errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	word, err := h.service.GetWord(id)
	if err != nil {
		errors := err

		response := helper.APIResponse("Failed to get word detail", http.StatusInternalServerError, nil, errors)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := dictionary.FormatDictionary(word)

	response := helper.APIResponse("Successfully to get word detail", http.StatusOK, data, nil)
	c.JSON(http.StatusOK, response)
}
