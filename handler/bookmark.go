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

func (h *bookmarkHandler) MarkedAndUnmarkedWord(c *gin.Context) {
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

	// check if user already bookmarked the word
	isMarked, err := h.service.FindByUserIDAndDictionaryID(input.User.ID, input.DictionaryID)
	if err != nil {
		response := helper.APIResponse("Failed to mark word", http.StatusInternalServerError, nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if isMarked {
		// if user already bookmarked the word, then unmark it
		err = h.service.UnmarkWord(input)
		if err != nil {
			response := helper.APIResponse("Failed to unmark word", http.StatusInternalServerError, nil, err.Error())
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		response := helper.APIResponse("Successfully to unmark word", http.StatusOK, nil, nil)
		c.JSON(http.StatusOK, response)
		return
	}

	newBookmark, err := h.service.MarkWord(input)
	if err != nil {
		response := helper.APIResponse("Failed to mark word", http.StatusInternalServerError, nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Successfully to mark word", http.StatusOK, newBookmark, nil)
	c.JSON(http.StatusOK, response)
}
