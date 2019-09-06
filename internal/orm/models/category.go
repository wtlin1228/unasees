package models

// Category defines a category for the app
type Category struct {
	BaseModelSoftDelete
	Name string `gorm:"not null;unique_index:idx_name"`

	Desserts []*Dessert
}
