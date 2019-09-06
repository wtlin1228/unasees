package models

import "github.com/gofrs/uuid"

// Dessert defines a dessert for the app
type Dessert struct {
	BaseModelSoftDelete
	Name           string
	Description    *string
	Unit           string
	Amount         int
	AmountMinimum  int
	AmountInterval int
	DegreeTop      string
	DegreeDown     string
	BakingTime     int // minutes
	BigImageURL    string
	SmallImageURL  string
	ThumbnailURL   string

	Steps            []*Step
	IngredientGroups []*IngredientGroup

	CategoryID uuid.UUID
}
