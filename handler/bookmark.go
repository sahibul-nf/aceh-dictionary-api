package handler

import (
	"aceh-dictionary-api/bookmark"
	"aceh-dictionary-api/helper"
	"aceh-dictionary-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookmarkHandler struct {
	service bookmark.Service
}

func NewBookmarkHandler(service bookmark.Service) *bookmarkHandler {
	return &bookmarkHandler{service}
}

func (h *bookmarkHandler) AddBookmark(c *gin.Context) {
	var input bookmark.MarkWordInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationFormat(err)

		response := helper.APIResponse("Failed to mark word", http.StatusBadRequest, nil, errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get user id from token
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	bookmark, err := h.service.MarkWord(input)
	if err != nil {
		response := helper.APIResponse("Failed to mark word", http.StatusInternalServerError, nil, err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Successfully to mark word", http.StatusOK, bookmark, nil)
	c.JSON(http.StatusOK, response)
}
