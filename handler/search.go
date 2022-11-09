package handler

import (
	"aceh-dictionary-api/helper"
	"aceh-dictionary-api/search"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type adviceHandler struct {
	service search.Service
}

func NewAdviceHandler(adviceService search.Service) *adviceHandler {
	return &adviceHandler{adviceService}
}

func (h *adviceHandler) GetAdvices(c *gin.Context) {

	input := c.Query("q")

	advices, err := h.service.GetRecommendation(input)
	if err != nil {
		response := helper.APIResponse("Failed to get word advices", http.StatusBadGateway, nil)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	if len(advices) < 1 {
		response := helper.APIResponse(fmt.Sprintf("Opps, no data found for similar to %s", input), http.StatusOK, advices)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Successfully to get word advice", http.StatusOK, advices)
	c.JSON(http.StatusOK, response)
}
