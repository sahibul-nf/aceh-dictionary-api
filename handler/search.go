package handler

import (
	"aceh-dictionary-api/helper"
	"aceh-dictionary-api/search"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type searchHandler struct {
	service search.Service
}

func NewSearchHandler(adviceService search.Service) *searchHandler {
	return &searchHandler{adviceService}
}

func (h *searchHandler) Search(c *gin.Context) {

	input := c.Query("q")

	recommendationList, err := h.service.GetRecommendationWords(input)
	if err != nil {
		errors := err

		response := helper.APIResponse("Failed to get recommendation list word", http.StatusInternalServerError, recommendationList, errors)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(recommendationList) < 1 {
		errors := errors.New("no recommendation list word found")

		response := helper.APIResponse(fmt.Sprintf("Opps, no data found for similar to %s", input), http.StatusNoContent, recommendationList, errors)
		c.JSON(http.StatusNoContent, response)
		return
	}

	response := helper.APIResponse("Successfully to get recommendation list word", http.StatusOK, recommendationList, nil)
	c.JSON(http.StatusOK, response)
}
