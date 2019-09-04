package resolvers

import (
	"context"

	gql "github.com/wtlin1228/go-gql-server/internal/gql/generated"
	models "github.com/wtlin1228/go-gql-server/internal/gql/models/generated"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePost(ctx context.Context, input models.PostInput) (*models.Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input models.PostInput) (*models.Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context, id *string) (*models.Users, error) {
	panic("not implemented")
}
func (r *queryResolver) Posts(ctx context.Context, id *string) (*models.Posts, error) {
	panic("not implemented")
}
