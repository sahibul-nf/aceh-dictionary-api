package bookmark

import "aceh-dictionary-api/user"

type MarkWordInput struct {
	DictionaryID int `json:"dictionary_id" binding:"required"`
	User         user.User
}
