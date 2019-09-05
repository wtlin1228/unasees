package resolvers

import (
	"context"

	"github.com/gofrs/uuid"
	gqlmodels "github.com/wtlin1228/go-gql-server/internal/gql/models"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// Mutations
func (r *mutationResolver) CreatePost(ctx context.Context, input gqlmodels.PostInput) (*models.Post, error) {
	return postCreateUpdate(r, input, false)
}
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input gqlmodels.PostInput) (*models.Post, error) {
	return postCreateUpdate(r, input, true, id)
}
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	return postDelete(r, id)
}

// Queries
func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	r.ORM.DB.Find(&posts)
	return posts, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*models.Post, error) {
	post := &models.Post{}
	r.ORM.DB.First(&post)
	return post, nil
}

type postResolver struct{ *Resolver }

func (r *postResolver) ID(ctx context.Context, obj *models.Post) (string, error) {
	return obj.ID.String(), nil
}
func (r *postResolver) User(ctx context.Context, obj *models.Post) (*models.User, error) {
	return r.Query().User(ctx, obj.UserID.String())
}

// Mutation Helper functions
func postCreateUpdate(r *mutationResolver, input gqlmodels.PostInput, update bool, ids ...string) (*models.Post, error) {
	dbo, err := GQLInputPostToDBPost(&input, update, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if !update {
		db = db.Create(dbo).First(dbo) // Create the post
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

func postDelete(r *mutationResolver, id string) (bool, error) {
	whereID := "id = ?"
	// Convert id from type string to type uuid.UUID
	convertedID, err := uuid.FromString(id)
	if err != nil {
		return false, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	// Find the post
	dbPost := &models.Post{}
	err = db.Where(whereID, convertedID).First(dbPost).Error
	if err != nil {
		return false, err
	}
	// Delete the post
	if err := db.Delete(dbPost).Error; err != nil {
		db.RollbackUnlessCommitted()
		return false, err
	}
	db = db.Commit()
	return true, nil
}

// GQLInputPostToDBPost transforms [post] gql input to db model
func GQLInputPostToDBPost(i *gqlmodels.PostInput, update bool, ids ...string) (o *models.Post, err error) {
	o = &models.Post{
		Title:   *i.Title,
		Content: i.Content,
	}
	// convert the id from type String to type uuid.UUID
	if len(ids) > 0 {
		updID, err := uuid.FromString(ids[0])
		if err != nil {
			return nil, err
		}
		o.ID = updID
	}
	return o, err
}
