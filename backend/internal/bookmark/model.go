package bookmark

import "github.com/google/uuid"

type BookmarkResponseModel struct {
	Id          uuid.UUID
	Title       string
	Description string
	Link        string
	Category    string
	Tags        []string
}
