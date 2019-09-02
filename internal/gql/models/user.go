package models

type User struct {
	ID     *string `json:"id"`
	Email  *string `json:"email"`
	UserID *string `json:"userId"`
}

type UserInput struct {
	Email  *string `json:"email"`
	UserID *string `json:"userId"`
}
