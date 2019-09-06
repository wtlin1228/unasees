package resolvers

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	gqlmodels "github.com/wtlin1228/go-gql-server/internal/gql/models"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// Mutations
func (r *mutationResolver) CreateIngredient(ctx context.Context, input *gqlmodels.IngredientInput) (*models.Ingredient, error) {
	return ingredientCreateUpdate(r, input, false)
}
func (r *mutationResolver) UpdateIngredient(ctx context.Context, id string, input *gqlmodels.IngredientInput) (*models.Ingredient, error) {
	return ingredientCreateUpdate(r, input, true, id)
}
func (r *mutationResolver) DeleteIngredient(ctx context.Context, id string) (bool, error) {
	return ingredientDelete(r, id)
}

// Queries
func (r *queryResolver) Ingredients(ctx context.Context) ([]*models.Ingredient, error) {
	var ingredients []*models.Ingredient
	r.ORM.DB.Find(&ingredients)
	return ingredients, nil
}
func (r *queryResolver) Ingredient(ctx context.Context, id string) (*models.Ingredient, error) {
	ingredient := &models.Ingredient{}
	r.ORM.DB.First(&ingredient)
	return ingredient, nil
}

// Ingredient resolvers
type ingredientResolver struct{ *Resolver }

func (r *ingredientResolver) ID(ctx context.Context, obj *models.Ingredient) (string, error) {
	return obj.ID.String(), nil
}
func (r *ingredientResolver) IngredientGroup(ctx context.Context, obj *models.Ingredient) (*models.IngredientGroup, error) {
	return r.Query().IngredientGroup(ctx, obj.IngredientGroupID.String())
}

// Mutation Helper functions
func ingredientCreateUpdate(r *mutationResolver, input *gqlmodels.IngredientInput, update bool, ids ...string) (*models.Ingredient, error) {
	dbo, err := gqlInputIngredientToDBIngredient(input, update, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if !update {
		db = db.Create(dbo).First(dbo) // Create the Ingredient
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

func ingredientDelete(r *mutationResolver, id string) (bool, error) {
	// Convert id from type string to type uuid.UUID
	convertedID, err := uuid.FromString(id)
	if err != nil {
		return false, err
	}
	// Find the Ingredient
	whereID := "id = ?"
	dbIngredient := &models.Ingredient{}
	err = r.ORM.DB.Where(whereID, convertedID).First(dbIngredient).Error
	if err != nil {
		return false, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if err := db.Delete(dbIngredient).Error; err != nil {
		db.RollbackUnlessCommitted()
		return false, err
	}
	db = db.Commit()
	return true, nil
}

func ingredientDeleteCascade(db *gorm.DB, ingredient *models.Ingredient) error {
	// Delete the Ingredient
	if err := db.Delete(ingredient).Error; err != nil {
		db.RollbackUnlessCommitted()
		return err
	}
	return nil
}

// gqlInputIngredientToDBIngredient transforms [Ingredient] gql input to db model
func gqlInputIngredientToDBIngredient(i *gqlmodels.IngredientInput, update bool, ids ...string) (o *models.Ingredient, err error) {
	o = &models.Ingredient{
		Name:   *i.Name,
		Unit:   *i.Unit,
		Amount: *i.Amount,
	}
	// Convert the IngredientGroupID from type String to type uuid.UUID
	parrentID, err := uuid.FromString(*i.IngredientGroupID)
	if err != nil {
		return nil, err
	}
	o.IngredientGroupID = parrentID
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
