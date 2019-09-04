package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var (
	uname                    = "Test User"
	fname                    = "Test"
	lname                    = "User"
	nname                    = "Foo Bar"
	description              = "This is the first user ever!"
	location                 = "His house, maybe? Wouldn't know"
	pcontent                 = "Test post Content"
	firstUser   *models.User = &models.User{
		Email:       "test@test.com",
		Name:        &uname,
		FirstName:   &fname,
		LastName:    &lname,
		NickName:    &nname,
		Description: &description,
		Location:    &location,
		Posts: []*models.Post{
			&models.Post{
				Title:   "Test post title",
				Content: &pcontent,
			},
		},
	}
)

// SeedUsers inserts the first users
var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstUser).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstUser).Error
	},
}
