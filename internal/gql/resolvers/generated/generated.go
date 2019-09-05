package resolvers

import (
	"context"

	"github.com/wtlin1228/go-gql-server/internal/gql/generated"
	gqlmodels "github.com/wtlin1228/go-gql-server/internal/gql/models"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Post() generated.PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input gqlmodels.UserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input gqlmodels.UserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePost(ctx context.Context, input gqlmodels.PostInput) (*models.Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input gqlmodels.PostInput) (*models.Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}

type postResolver struct{ *Resolver }

func (r *postResolver) ID(ctx context.Context, obj *models.Post) (string, error) {
	panic("not implemented")
}
func (r *postResolver) User(ctx context.Context, obj *models.Post) (*models.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Post(ctx context.Context, id string) (*models.Post, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	panic("not implemented")
}
