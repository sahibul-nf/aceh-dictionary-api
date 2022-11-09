package handler

import (
	"aceh-dictionary-api/helper"
	"aceh-dictionary-api/search"
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
		response := helper.APIResponse("Failed to get recommendation list word", http.StatusBadGateway, nil)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	if len(recommendationList) < 1 {
		response := helper.APIResponse(fmt.Sprintf("Opps, no data found for similar to %s", input), http.StatusOK, recommendationList)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Successfully to get recommendation list word", http.StatusOK, recommendationList)
	c.JSON(http.StatusOK, response)
}
