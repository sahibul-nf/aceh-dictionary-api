package handler

import (
	"aceh-dictionary-api/advice"
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
		c.JSON(http.StatusBadRequest, advice.Advice{})
		return
	}

	advices, err := h.service.GetAdvices(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, advice.Advice{})
		return
	}

	c.JSON(http.StatusOK, advices)
}
