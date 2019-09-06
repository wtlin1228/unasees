package resolvers

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	gqlmodels "github.com/wtlin1228/go-gql-server/internal/gql/models"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// Mutations
func (r *mutationResolver) CreateStep(ctx context.Context, input *gqlmodels.StepInput) (*models.Step, error) {
	return stepCreateUpdate(r, input, false)
}
func (r *mutationResolver) UpdateStep(ctx context.Context, id string, input *gqlmodels.StepInput) (*models.Step, error) {
	return stepCreateUpdate(r, input, true, id)
}
func (r *mutationResolver) DeleteStep(ctx context.Context, id string) (bool, error) {
	return stepDelete(r, id)
}

// Queries
func (r *queryResolver) Steps(ctx context.Context) ([]*models.Step, error) {
	var steps []*models.Step
	r.ORM.DB.Find(&steps)
	return steps, nil
}
func (r *queryResolver) Step(ctx context.Context, id string) (*models.Step, error) {
	step := &models.Step{}
	r.ORM.DB.First(&step)
	return step, nil
}

// Step resolvers
type stepResolver struct{ *Resolver }

func (r *stepResolver) ID(ctx context.Context, obj *models.Step) (string, error) {
	return obj.ID.String(), nil
}
func (r *stepResolver) Dessert(ctx context.Context, obj *models.Step) (*models.Dessert, error) {
	return r.Query().Dessert(ctx, obj.DessertID.String())
}

// Mutation Helper functions
func stepCreateUpdate(r *mutationResolver, input *gqlmodels.StepInput, update bool, ids ...string) (*models.Step, error) {
	dbo, err := gqlInputStepToDBStep(input, update, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if !update {
		db = db.Create(dbo).First(dbo) // Create the Step
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

func stepDelete(r *mutationResolver, id string) (bool, error) {
	// Convert id from type string to type uuid.UUID
	convertedID, err := uuid.FromString(id)
	if err != nil {
		return false, err
	}
	// Find the Step
	whereID := "id = ?"
	dbStep := &models.Step{}
	err = r.ORM.DB.Where(whereID, convertedID).First(dbStep).Error
	if err != nil {
		return false, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if err := db.Delete(dbStep).Error; err != nil {
		db.RollbackUnlessCommitted()
		return false, err
	}
	db = db.Commit()
	return true, nil
}

func stepDeleteCascade(db *gorm.DB, step *models.Step) error {
	// Delete the Step
	if err := db.Delete(step).Error; err != nil {
		db.RollbackUnlessCommitted()
		return err
	}
	return nil
}

// gqlInputStepToDBStep transforms [Step] gql input to db model
func gqlInputStepToDBStep(i *gqlmodels.StepInput, update bool, ids ...string) (o *models.Step, err error) {
	o = &models.Step{
		Name:    *i.Name,
		Content: *i.Content,
		Notice:  *i.Notice,
		Order:   *i.Order,
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
