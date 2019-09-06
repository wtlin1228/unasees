package models

import "github.com/gofrs/uuid"

// IngredientGroup defines a ingredientGroup for the app
type IngredientGroup struct {
	BaseModelSoftDelete
	Name string

	Ingredients []*Ingredient

	DessertID uuid.UUID
}
