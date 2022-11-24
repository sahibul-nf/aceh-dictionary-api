package handler

import (
	"aceh-dictionary-api/bookmark"
	"aceh-dictionary-api/helper"
	"aceh-dictionary-api/user"
	"net/http"
	"strconv"

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

		response := helper.APIResponse("Failed to mark or unmark word", http.StatusBadRequest, nil, errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get user id from token
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// check if user already bookmarked the word
	markedWord, err := h.service.FindByUserIDAndDictionaryID(input.User.ID, input.DictionaryID)
	if err != nil {
		response := helper.APIResponse("Failed to mark or unmark word", http.StatusInternalServerError, nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	isMarked := markedWord.DictionaryID != 0

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

func (h *bookmarkHandler) GetMarkedWords(c *gin.Context) {
	param := c.Query("dictionary_id")
	dictionaryID, _ := strconv.Atoi(param)

	// Get user id from token
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	markedWords, err := h.service.FindByUserIDAndDictionaryID(userID, dictionaryID)
	if err != nil {
		response := helper.APIResponse("Failed to get marked word", http.StatusInternalServerError, nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if markedWords.DictionaryID == 0 {
		response := helper.APIResponse("No marked word", http.StatusOK, nil, nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Successfully to get marked word", http.StatusOK, markedWords, nil)
	c.JSON(http.StatusOK, response)
}

func (h *bookmarkHandler) GetMarkedWordsByUserID(c *gin.Context) {
	// Get user id from token
	currentUser := c.MustGet("currentUser").(user.User)

	markedWords, err := h.service.FindByUserID(currentUser.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get marked words", http.StatusInternalServerError, nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Successfully to get marked words", http.StatusOK, markedWords, nil)
	c.JSON(http.StatusOK, response)
}

func (h *bookmarkHandler) DeleteMarkedWord(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	// Get user id from token
	currentUser := c.MustGet("currentUser").(user.User)

	bookmark, err := h.service.FindByID(id)
	if err != nil {
		response := helper.APIResponse("Failed to delete marked word", http.StatusInternalServerError, nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if bookmark.ID == 0 {
		response := helper.APIResponse("Failed to delete marked word", http.StatusNotFound, nil, "Marked word not found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	if bookmark.UserID != currentUser.ID {
		response := helper.APIResponse("Failed to delete marked word", http.StatusForbidden, nil, "You are not allowed to delete this marked word, because you are not the owner")
		c.JSON(http.StatusForbidden, response)
		return
	}

	err = h.service.DeleteBookmarkItemByUserID(id, currentUser.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete marked word", http.StatusInternalServerError, nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Successfully to delete marked word", http.StatusOK, nil, nil)
	c.JSON(http.StatusOK, response)
}

func (h *bookmarkHandler) DeleteAllMarkedWordsByUserID(c *gin.Context) {
	// Get user id from token
	currentUser := c.MustGet("currentUser").(user.User)

	err := h.service.DeleteAllBookmarkByUserID(currentUser.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete all marked words", http.StatusInternalServerError, nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Successfully to delete all marked words", http.StatusOK, nil, nil)
	c.JSON(http.StatusOK, response)
}
