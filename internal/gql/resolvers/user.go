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

// CreateUser creates a record
func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	return userCreateUpdate(r, input, false)
}

// UpdateUser updates a record
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UserInput) (*models.User, error) {
	return userCreateUpdate(r, input, true, id)
}

// DeleteUser deletes a record
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	return userDelete(r, id)
}

// Users lists records
func (r *queryResolver) Users(ctx context.Context, id *string) (*models.Users, error) {
	return userList(r, id)
}

// ## Helper functions

func userCreateUpdate(r *mutationResolver, input models.UserInput, update bool, ids ...string) (*models.User, error) {
	dbo, err := tf.GQLInputUserToDBUser(&input, update, ids...)
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
	gql, err := tf.DBUserToGQLUser(dbo)
	if err != nil {
		db.RollbackUnlessCommitted()
		return nil, err
	}
	db = db.Commit()
	return gql, db.Error
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
	dbUser := &dbm.User{}
	err = db.Where(whereID, convertedID).First(dbUser).Error
	if err != nil {
		return false, err
	}
	// Find the user's posts
	dbPosts := []*dbm.Post{}
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

func userList(r *queryResolver, id *string) (*models.Users, error) {
	entity := "users"
	whereID := "id = ?"
	record := &models.Users{}
	dbRecords := []*dbm.User{}
	db := r.ORM.DB.New()
	if id != nil {
		db = db.Where(whereID, *id)
	}
	db = db.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		if rec, err := tf.DBUserToGQLUser(dbRec); err != nil {
			log.Println(entity, err)
		} else {
			// get user's posts
			dbPosts := []*dbm.Post{}
			db.Model(&dbRec).Related(&dbPosts, "Posts")
			for _, dbPost := range dbPosts {
				if gqlPost, err := tf.DBPostToGQLPost(dbPost); err != nil {
					log.Println("posts", err)
				} else {
					rec.Posts = append(rec.Posts, gqlPost)
				}
			}
			record.List = append(record.List, rec)
		}
	}
	return record, db.Error
}
