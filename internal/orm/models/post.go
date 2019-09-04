package models

import "github.com/gofrs/uuid"

// Post defines a post for the app
type Post struct {
	BaseModelSoftDelete
	Title   string
	Content *string
	UserID  uuid.UUID // User has many Posts
}
