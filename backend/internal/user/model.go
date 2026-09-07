package user

import "github.com/google/uuid"

type ProfileResponseModel struct {
	Id        uuid.UUID
	Email     string
	UserName  string
	FirstName string
	LastName  string
	TagLine   string
}
