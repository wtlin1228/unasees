package models

import "github.com/gofrs/uuid"

// Ingredient defines a ingredient for the app
type Ingredient struct {
	BaseModelSoftDelete
	Name   string
	Unit   string
	Amount int

	IngredientGroupID uuid.UUID
}
