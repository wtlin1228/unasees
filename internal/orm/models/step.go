package models

import "github.com/gofrs/uuid"

// Step defines a step for the app
type Step struct {
	BaseModelSoftDelete
	Name    string
	Content string
	Notice  string
	Order   int

	DessertID uuid.UUID
}
