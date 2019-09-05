package resolvers

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	gqlmodels "github.com/wtlin1228/go-gql-server/internal/gql/models"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// Mutations
func (r *mutationResolver) CreateUser(ctx context.Context, input gqlmodels.UserInput) (*models.User, error) {
	return userCreateUpdate(r, input, false)
}
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input gqlmodels.UserInput) (*models.User, error) {
	return userCreateUpdate(r, input, true, id)
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	return userDelete(r, id)
}

// Queries
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	r.ORM.DB.Preload("Posts").Find(&users)
	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	r.ORM.DB.Preload("Posts").First(&user)
	return user, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return obj.ID.String(), nil
}

// Mutation Helper functions
func userCreateUpdate(r *mutationResolver, input gqlmodels.UserInput, update bool, ids ...string) (*models.User, error) {
	dbo, err := GQLInputUserToDBUser(&input, update, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if !update {
		db = db.Create(dbo).First(dbo) // Create the user
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

func userDelete(r *mutationResolver, id string) (bool, error) {
	whereID := "id = ?"
	// Convert id to uuid.UUID from string
	convertedID, err := uuid.FromString(id)
	if err != nil {
		return false, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	// Find the user
	dbUser := &models.User{}
	err = db.Where(whereID, convertedID).First(dbUser).Error
	if err != nil {
		return false, err
	}
	// Find the user's posts
	dbPosts := []*models.Post{}
	db.Model(&dbUser).Related(&dbPosts, "Posts")
	// Delete posts
	for _, dbPost := range dbPosts {
		if err := db.Delete(dbPost).Error; err != nil {
			db.RollbackUnlessCommitted()
			return false, err
		}
	}
	// Delete the user
	if err := db.Delete(dbUser).Error; err != nil {
		db.RollbackUnlessCommitted()
		return false, err
	}
	db = db.Commit()
	return true, nil
}

// GQLInputUserToDBUser transforms [user] gql input to db model
func GQLInputUserToDBUser(i *gqlmodels.UserInput, update bool, ids ...string) (o *models.User, err error) {
	o = &models.User{
		UserID:      i.UserID,
		Name:        i.Name,
		FirstName:   i.FirstName,
		LastName:    i.LastName,
		NickName:    i.NickName,
		Description: i.Description,
		Location:    i.Location,
	}
	if i.Email == nil && !update {
		return nil, errors.New("Field [email] is required")
	}
	if i.Email != nil {
		o.Email = *i.Email
	}
	if len(ids) > 0 {
		updID, err := uuid.FromString(ids[0])
		if err != nil {
			return nil, err
		}
		o.ID = updID
	}
	return o, err
}
