package resolvers

import (
	"context"
	"log"

	"github.com/gofrs/uuid"
	models "github.com/wtlin1228/go-gql-server/internal/gql/models"
	tf "github.com/wtlin1228/go-gql-server/internal/gql/resolvers/transformations"
	dbm "github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func (r *mutationResolver) CreatePost(ctx context.Context, input models.PostInput) (*models.Post, error) {
	return postCreateUpdate(r, input, false)
}
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input models.PostInput) (*models.Post, error) {
	return postCreateUpdate(r, input, true, id)
}
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	return postDelete(r, id)
}

func (r *queryResolver) Posts(ctx context.Context, id *string) (*models.Posts, error) {
	return postList(r, id)
}

// ## Helper functions

func postCreateUpdate(r *mutationResolver, input models.PostInput, update bool, ids ...string) (*models.Post, error) {
	dbo, err := tf.GQLInputPostToDBPost(&input, update, ids...)
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
	gql, err := tf.DBPostToGQLPost(dbo)
	if err != nil {
		db.RollbackUnlessCommitted()
		return nil, err
	}
	db = db.Commit()
	return gql, db.Error
}

func postDelete(r *mutationResolver, id string) (bool, error) {
	whereID := "id = ?"
	// Convert id to uuid.UUID from string
	convertedID, err := uuid.FromString(id)
	if err != nil {
		return false, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	// Find the post
	dbPost := &dbm.Post{}
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

func postList(r *queryResolver, id *string) (*models.Posts, error) {
	entity := "posts"
	whereID := "id = ?"
	record := &models.Posts{}
	dbRecords := []*dbm.Post{}
	db := r.ORM.DB.New()
	if id != nil {
		db = db.Where(whereID, *id)
	}
	db = db.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		if rec, err := tf.DBPostToGQLPost(dbRec); err != nil {
			log.Println(entity, err)
		} else {
			// get the post's owner
			dbUser := &dbm.User{}
			db.Where(whereID, dbRec.UserID).First(dbUser)
			if gqlUser, err := tf.DBUserToGQLUser(dbUser); err != nil {
				log.Println(entity, err)
			} else {
				rec.User = gqlUser
			}

			record.List = append(record.List, rec)
		}
	}
	return record, db.Error
}
