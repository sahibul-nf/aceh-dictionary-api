package handler

import (
	"aceh-dictionary-api/helper"
	"aceh-dictionary-api/vocabulary"
	"net/http"

	"github.com/gin-gonic/gin"
)

type vocabularyHandler struct {
	service vocabulary.Service
}

func NewVocabularyHandler(vocabService vocabulary.Service) *vocabularyHandler {
	return &vocabularyHandler{vocabService}
}

func (h *vocabularyHandler) AddNewVocabulary(c *gin.Context) {
	var input vocabulary.VocabularyInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("error", "Failed to add new vocabulary", http.StatusBadRequest, false)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isSuccess, err := h.service.SaveData(input)
	if err != nil {
		response := helper.APIResponse("error", "Failed to add new vocabulary", http.StatusBadRequest, false)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !isSuccess {
		response := helper.APIResponse("error", "Failed to add new vocabulary", http.StatusBadRequest, false)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success", "Successfully to add new vocabulary", http.StatusOK, isSuccess)
	c.JSON(http.StatusOK, response)
}

func (h *vocabularyHandler) GetAllVocabularyData(c *gin.Context) {

	data, err := h.service.GetAllData()
	if err != nil {
		response := helper.APIResponse("error", "Failed to get all vocabulary", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success", "Successfully to get all vocabulary", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}
