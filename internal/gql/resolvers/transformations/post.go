package transformations

import (
	"github.com/gofrs/uuid"
	gql "github.com/wtlin1228/go-gql-server/internal/gql/models"
	dbm "github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// DBPostToGQLPost transforms [post] db input to gql type
func DBPostToGQLPost(i *dbm.Post) (o *gql.Post, err error) {
	o = &gql.Post{
		ID:        i.ID.String(),
		Title:     i.Title,
		Content:   i.Content,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
	return o, err
}

// GQLInputPostToDBPost transforms [post] gql input to db model
func GQLInputPostToDBPost(i *gql.PostInput, update bool, ids ...string) (o *dbm.Post, err error) {
	o = &dbm.Post{
		Title:   *i.Title,
		Content: i.Content,
	}
	updUserID, err := uuid.FromString(*i.UserID)
	if err != nil {
		return nil, err
	}
	o.UserID = updUserID

	if len(ids) > 0 {
		updID, err := uuid.FromString(ids[0])
		if err != nil {
			return nil, err
		}
		o.ID = updID
	}
	return o, err
}
