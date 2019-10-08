package resolvers

import (
	"context"

	gqlmodels "github.com/wtlin1228/go-gql-server/internal/gql/models"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// Mutations
func (r *mutationResolver) Login(ctx context.Context, input *gqlmodels.UserInput) (string, error) {
	return login(r, input)
}

// Queries
func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	panic("not implemented")
}

// Step resolvers
type userResolver struct{ *Resolver }

func login(r *mutationResolver, input *gqlmodels.UserInput) (string, error) {
	// Implement your login logic here
	return "MyFakeToken", nil
}
