package models

// User defines a user for the app
type User struct {
	BaseModelSoftDelete
	Username string `gorm:"not null;unique_index:idx_name"`
	Password string `gorm:"not null"`
}
