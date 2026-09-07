package api

import "github.com/guregu/null/v6"

type ProfileUpdateRequest struct {
	FirstName null.String `json:"first_name,omitempty"`
	LastName  null.String `json:"last_name,omitempty"`
	TagLine   null.String `json:"tag_line,omitempty"`
}

type ProfileResponse struct {
	Email     string `json:"email"`
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TagLine   string `json:"tag_line"`
}
