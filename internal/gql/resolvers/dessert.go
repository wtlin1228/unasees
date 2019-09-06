package resolvers

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	gqlmodels "github.com/wtlin1228/go-gql-server/internal/gql/models"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// Mutations
func (r *mutationResolver) CreateDessert(ctx context.Context, input *gqlmodels.DessertInput) (*models.Dessert, error) {
	return dessertCreateUpdate(r, input, false)
}
func (r *mutationResolver) UpdateDessert(ctx context.Context, id string, input *gqlmodels.DessertInput) (*models.Dessert, error) {
	return dessertCreateUpdate(r, input, true, id)
}
func (r *mutationResolver) DeleteDessert(ctx context.Context, id string) (bool, error) {
	return dessertDelete(r, id)
}

// Queries
func (r *queryResolver) Desserts(ctx context.Context) ([]*models.Dessert, error) {
	var desserts []*models.Dessert
	r.ORM.DB.Find(&desserts)
	return desserts, nil
}
func (r *queryResolver) Dessert(ctx context.Context, id string) (*models.Dessert, error) {
	dessert := &models.Dessert{}
	r.ORM.DB.First(&dessert)
	return dessert, nil
}

// Dessert resolvers
type dessertResolver struct{ *Resolver }

func (r *dessertResolver) ID(ctx context.Context, obj *models.Dessert) (string, error) {
	return obj.ID.String(), nil
}
func (r *dessertResolver) StepList(ctx context.Context, obj *models.Dessert) ([]*models.Step, error) {
	return getStepsOfDessert(r.ORM.DB, obj), nil
}
func (r *dessertResolver) IngredientGroupList(ctx context.Context, obj *models.Dessert) ([]*models.IngredientGroup, error) {
	return getIngredientGroupsOfDessert(r.ORM.DB, obj), nil
}
func (r *dessertResolver) Category(ctx context.Context, obj *models.Dessert) (*models.Category, error) {
	return r.Query().Category(ctx, obj.CategoryID.String())
}

// Mutation Helper functions
func dessertCreateUpdate(r *mutationResolver, input *gqlmodels.DessertInput, update bool, ids ...string) (*models.Dessert, error) {
	dbo, err := gqlInputDessertToDBDessert(input, update, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if !update {
		db = db.Create(dbo).First(dbo) // Create the Dessert
	} else {
		db = db.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	if db.Error != nil {
		db.RollbackUnlessCommitted()
		return nil, db.Error
	}
	db = db.Commit()
	return dbo, nil
}

func dessertDelete(r *mutationResolver, id string) (bool, error) {
	// Convert id from type string to type uuid.UUID
	convertedID, err := uuid.FromString(id)
	if err != nil {
		return false, err
	}
	// Find the Dessert
	whereID := "id = ?"
	dbDessert := &models.Dessert{}
	err = r.ORM.DB.Where(whereID, convertedID).First(dbDessert).Error
	if err != nil {
		return false, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if err := dessertDeleteCascade(db, dbDessert); err != nil {
		return false, err
	}
	db = db.Commit()
	return true, nil
}

func dessertDeleteCascade(db *gorm.DB, dessert *models.Dessert) error {
	// Delete Steps
	steps := getStepsOfDessert(db, dessert)
	for _, step := range steps {
		if err := stepDeleteCascade(db, step); err != nil {
			return err
		}
	}
	// Delete IngredientGroups
	ingredientGroups := getIngredientGroupsOfDessert(db, dessert)
	for _, ingredient := range ingredientGroups {
		if err := ingredientGroupDeleteCascade(db, ingredient); err != nil {
			return err
		}
	}
	// Delete the Dessert
	if err := db.Delete(dessert).Error; err != nil {
		db.RollbackUnlessCommitted()
		return err
	}
	return nil
}

func getStepsOfDessert(db *gorm.DB, dessert *models.Dessert) []*models.Step {
	var steps []*models.Step
	db.Model(&dessert).Related(&steps, "Steps")
	return steps
}

func getIngredientGroupsOfDessert(db *gorm.DB, dessert *models.Dessert) []*models.IngredientGroup {
	var ingredientGroups []*models.IngredientGroup
	db.Model(&dessert).Related(&ingredientGroups, "ingredientGroups")
	return ingredientGroups
}

// gqlInputDessertToDBDessert transforms [Dessert] gql input to db model
func gqlInputDessertToDBDessert(i *gqlmodels.DessertInput, update bool, ids ...string) (o *models.Dessert, err error) {
	o = &models.Dessert{
		Name:           *i.Name,
		Description:    i.Description,
		Unit:           *i.Unit,
		Amount:         *i.Amount,
		AmountMinimum:  *i.AmountMinimum,
		AmountInterval: *i.AmountInterval,
		DegreeTop:      *i.DegreeTop,
		DegreeDown:     *i.DegreeDown,
		BakingTime:     *i.BakingTime,
		BigImageURL:    *i.BigImageURL,
		SmallImageURL:  *i.SmallImageURL,
		ThumbnailURL:   *i.ThumbnailURL,
	}
	// Convert the CategoryID from type String to type uuid.UUID
	parrentID, err := uuid.FromString(*i.CategoryID)
	if err != nil {
		return nil, err
	}
	o.CategoryID = parrentID
	// Convert the id from type String to type uuid.UUID
	if len(ids) > 0 {
		updID, err := uuid.FromString(ids[0])
		if err != nil {
			return nil, err
		}
		o.ID = updID
	}
	return o, err
}
