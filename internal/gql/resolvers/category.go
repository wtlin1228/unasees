package resolvers

import (
	"context"

	"github.com/gofrs/uuid"
	gqlmodels "github.com/wtlin1228/go-gql-server/internal/gql/models"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// Mutations
func (r *mutationResolver) CreateCategory(ctx context.Context, input *gqlmodels.CategoryInput) (*models.Category, error) {
	return categoryCreateUpdate(r, input, false)
}
func (r *mutationResolver) UpdateCategory(ctx context.Context, id string, input *gqlmodels.CategoryInput) (*models.Category, error) {
	return categoryCreateUpdate(r, input, true, id)
}
func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (bool, error) {
	return categoryDelete(r, id)
}

// Queries
func (r *queryResolver) Categories(ctx context.Context) ([]*models.Category, error) {
	var categories []*models.Category
	r.ORM.DB.Find(&categories)
	return categories, nil
}
func (r *queryResolver) Category(ctx context.Context, id string) (*models.Category, error) {
	category := &models.Category{}
	r.ORM.DB.First(&category)
	return category, nil
}

// Category resolvers
type categoryResolver struct{ *Resolver }

func (r *categoryResolver) ID(ctx context.Context, obj *models.Category) (string, error) {
	return obj.ID.String(), nil
}
func (r *categoryResolver) DessertList(ctx context.Context, obj *models.Category) ([]*models.Dessert, error) {
	var desserts []*models.Dessert
	r.ORM.DB.Model(&obj).Related(&desserts, "Desserts")
	return desserts, nil
}

// Mutation Helper functions
func categoryCreateUpdate(r *mutationResolver, input *gqlmodels.CategoryInput, update bool, ids ...string) (*models.Category, error) {
	dbo, err := gqlInputCategoryToDBCategory(input, update, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if !update {
		db = db.Create(dbo).First(dbo) // Create the category
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

func categoryDelete(r *mutationResolver, id string) (bool, error) {
	whereID := "id = ?"
	// Convert id from type string to type uuid.UUID
	convertedID, err := uuid.FromString(id)
	if err != nil {
		return false, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	// Find the category
	dbCategory := &models.Category{}
	err = db.Where(whereID, convertedID).First(dbCategory).Error
	if err != nil {
		return false, err
	}
	// Delete the category
	if err := db.Delete(dbCategory).Error; err != nil {
		db.RollbackUnlessCommitted()
		return false, err
	}
	db = db.Commit()
	return true, nil
}

// gqlInputCategoryToDBCategory transforms [category] gql input to db model
func gqlInputCategoryToDBCategory(i *gqlmodels.CategoryInput, update bool, ids ...string) (o *models.Category, err error) {
	o = &models.Category{
		Name: *i.Name,
	}
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
