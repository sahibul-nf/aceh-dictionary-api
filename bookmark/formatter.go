package bookmark

import "aceh-dictionary-api/dictionary"

type BookmarkFormatter struct {
	ID           int `json:"id"`
	UserID       int `json:"user_id"`
	DictionaryID int `json:"dictionary_id"`
}

func FormatBookmark(bookmark Bookmark) BookmarkFormatter {
	bookmarkFormatter := BookmarkFormatter{
		ID:           bookmark.ID,
		UserID:       bookmark.UserID,
		DictionaryID: bookmark.DictionaryID,
	}

	return bookmarkFormatter
}

type BookmarksFormatter struct {
	ID         int                            `json:"id"`
	UserID     int                            `json:"user_id"`
	Dictionary dictionary.DictionaryFormatter `json:"dictionary"`
}

func FormatBookmarks(bookmarks []Bookmark) []BookmarksFormatter {
	var bookmarksFormatter []BookmarksFormatter

	for _, bookmark := range bookmarks {
		bookmarkFormatter := BookmarksFormatter{
			ID:         bookmark.ID,
			UserID:     bookmark.UserID,
			Dictionary: dictionary.FormatDictionary(bookmark.Dictionary),
		}

		bookmarksFormatter = append(bookmarksFormatter, bookmarkFormatter)
	}

	return bookmarksFormatter
}
