package models

import (
	"time"
)

type Post struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Content   *string    `json:"content"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	User      *User      `json:"user"`
}

type PostInput struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
	UserID  *string `json:"userId"`
}

type Posts struct {
	Count *int    `json:"count"`
	List  []*Post `json:"list"`
}
