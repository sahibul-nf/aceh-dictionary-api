package handler

import (
	"aceh-dictionary-api/advice"
	"aceh-dictionary-api/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type adviceHandler struct {
	service advice.Service
}

func NewAdviceHandler(adviceService advice.Service) *adviceHandler {
	return &adviceHandler{adviceService}
}

func (h *adviceHandler) GetAdvices(c *gin.Context) {
	var input advice.QueryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("error", "Failed to get word advices", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	advices, err := h.service.GetAdvices(input)
	if err != nil {
		response := helper.APIResponse("error", "Failed to get word advices", http.StatusBadGateway, nil)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	if len(advices) < 1 {
		response := helper.APIResponse("success", fmt.Sprintf("Opps, no data found for similar to %s", input), http.StatusOK, advices)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("success", "Successfully to get word advice", http.StatusOK, advices)
	c.JSON(http.StatusOK, response)
}
