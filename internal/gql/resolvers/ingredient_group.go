package resolvers

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	gqlmodels "github.com/wtlin1228/go-gql-server/internal/gql/models"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// Mutations
func (r *mutationResolver) CreateIngredientGroup(ctx context.Context, input *gqlmodels.IngredientGroupInput) (*models.IngredientGroup, error) {
	return ingredientGroupCreateUpdate(r, input, false)
}
func (r *mutationResolver) UpdateIngredientGroup(ctx context.Context, id string, input *gqlmodels.IngredientGroupInput) (*models.IngredientGroup, error) {
	return ingredientGroupCreateUpdate(r, input, true, id)
}
func (r *mutationResolver) DeleteIngredientGroup(ctx context.Context, id string) (bool, error) {
	return ingredientGroupDelete(r, id)
}

// Queries
func (r *queryResolver) IngredientGroups(ctx context.Context) ([]*models.IngredientGroup, error) {
	var ingredientGroups []*models.IngredientGroup
	r.ORM.DB.Preload("Ingredients").Find(&ingredientGroups)
	return ingredientGroups, nil
}
func (r *queryResolver) IngredientGroup(ctx context.Context, id string) (*models.IngredientGroup, error) {
	ingredientGroup := &models.IngredientGroup{}
	r.ORM.DB.Preload("Ingredients").First(&ingredientGroup)
	return ingredientGroup, nil
}

// Ingredient Group resolvers
type ingredientGroupResolver struct{ *Resolver }

func (r *ingredientGroupResolver) ID(ctx context.Context, obj *models.IngredientGroup) (string, error) {
	return obj.ID.String(), nil
}
func (r *ingredientGroupResolver) IngredientList(ctx context.Context, obj *models.IngredientGroup) ([]*models.Ingredient, error) {
	return getIngredientsOfIngredientGroup(r.ORM.DB, obj), nil
}
func (r *ingredientGroupResolver) Dessert(ctx context.Context, obj *models.IngredientGroup) (*models.Dessert, error) {
	return r.Query().Dessert(ctx, obj.DessertID.String())
}

// Mutation Helper functions
func ingredientGroupCreateUpdate(r *mutationResolver, input *gqlmodels.IngredientGroupInput, update bool, ids ...string) (*models.IngredientGroup, error) {
	dbo, err := gqlInputIngredientGroupToDBIngredientGroup(input, update, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if !update {
		db = db.Create(dbo).First(dbo) // Create the IngredientGroup
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

func ingredientGroupDelete(r *mutationResolver, id string) (bool, error) {
	// Convert id from type string to type uuid.UUID
	convertedID, err := uuid.FromString(id)
	if err != nil {
		return false, err
	}
	// Find the IngredientGroup
	whereID := "id = ?"
	dbIngredientGroup := &models.IngredientGroup{}
	err = r.ORM.DB.Where(whereID, convertedID).First(dbIngredientGroup).Error
	if err != nil {
		return false, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if err := ingredientGroupDeleteCascade(db, dbIngredientGroup); err != nil {
		db.RollbackUnlessCommitted()
		return false, err
	}
	db = db.Commit()
	return true, nil
}

func ingredientGroupDeleteCascade(db *gorm.DB, ingredientGroup *models.IngredientGroup) error {
	// Delete Ingredients
	ingredients := getIngredientsOfIngredientGroup(db, ingredientGroup)
	for _, ingredient := range ingredients {
		if err := ingredientDeleteCascade(db, ingredient); err != nil {
			return err
		}
	}
	// Delete the IngredientGroup
	if err := db.Delete(ingredientGroup).Error; err != nil {
		db.RollbackUnlessCommitted()
		return err
	}
	return nil
}

func getIngredientsOfIngredientGroup(db *gorm.DB, ingredientGroup *models.IngredientGroup) []*models.Ingredient {
	var ingredients []*models.Ingredient
	db.Model(&ingredientGroup).Related(&ingredients, "Ingredients")
	return ingredients
}

// gqlInputIngredientGroupToDBIngredientGroup transforms [IngredientGroup] gql input to db model
func gqlInputIngredientGroupToDBIngredientGroup(i *gqlmodels.IngredientGroupInput, update bool, ids ...string) (o *models.IngredientGroup, err error) {
	o = &models.IngredientGroup{
		Name: *i.Name,
	}
	// Convert the DessertID from type String to type uuid.UUID
	parrentID, err := uuid.FromString(*i.DessertID)
	if err != nil {
		return nil, err
	}
	o.DessertID = parrentID
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
