package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var (
	cdescription                   = "This is my first test cake"
	firstCategory *models.Category = &models.Category{
		Name: "Cake",
		Desserts: []*models.Dessert{
			&models.Dessert{
				Name:           "Test Cake",
				Description:    &cdescription,
				Unit:           "個",
				Amount:         1,
				AmountMinimum:  1,
				AmountInterval: 1,
				DegreeTop:      "180",
				DegreeDown:     "180",
				BakingTime:     90,
				BigImageURL:    "http://static.food2fork.com/CadburyEgg1of1adf2.jpg",
				SmallImageURL:  "http://static.food2fork.com/CadburyEgg1of1adf2.jpg",
				ThumbnailURL:   "http://static.food2fork.com/CadburyEgg1of1adf2.jpg",
				Steps: []*models.Step{
					&models.Step{
						Name:    "This is step 1",
						Content: "Test Test Test Test Test Test Test Test Test Test Test Test Test Test Test Test Test Test Test Test",
						Notice:  "",
						Order:   0,
					},
					&models.Step{
						Name:    "This is step 2",
						Content: "Test Test Test Test Test Test Test Test Test Test Test Test Test Test Test Test Test Test Test Test",
						Notice:  "",
						Order:   1,
					},
				},
				IngredientGroups: []*models.IngredientGroup{
					&models.IngredientGroup{
						Name: "Test group 1",
						Ingredients: []*models.Ingredient{
							&models.Ingredient{
								Name:   "Test Ingredient 1",
								Unit:   "g",
								Amount: 100,
							},
							&models.Ingredient{
								Name:   "Test Ingredient 2",
								Unit:   "g",
								Amount: 150,
							},
						},
					},
					&models.IngredientGroup{
						Name: "Test group 2",
						Ingredients: []*models.Ingredient{
							&models.Ingredient{
								Name:   "Test Ingredient 3",
								Unit:   "g",
								Amount: 30,
							},
							&models.Ingredient{
								Name:   "Test Ingredient 4",
								Unit:   "g",
								Amount: 52,
							},
						},
					},
				},
			},
		},
	}
)

// SeedCategories inserts the first users
var SeedCategories *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_CATEGORIES",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstCategory).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstCategory).Error
	},
}
