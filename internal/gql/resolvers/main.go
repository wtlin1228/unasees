package resolvers

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/wtlin1228/go-gql-server/internal/errors"
	"github.com/wtlin1228/go-gql-server/internal/gql/generated"
	"github.com/wtlin1228/go-gql-server/internal/orm"
)

type contextKey string

var (
	UserIDCtxKey = contextKey("userID")
)

type Resolver struct {
	ORM *orm.ORM
}

func NewRootResolvers(orm *orm.ORM) generated.Config {
	c := generated.Config{
		Resolvers: &Resolver{
			ORM: orm, // pass in the ORM instance in the resolvers to be used
		},
	}

	// Schema Directive
	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		ctxUserID := ctx.Value(UserIDCtxKey)
		if ctxUserID == nil {
			return nil, errors.UnAuthorizedError
		}
		return next(ctx)
	}

	return c
}

func (r *Resolver) Category() generated.CategoryResolver {
	return &categoryResolver{r}
}
func (r *Resolver) Dessert() generated.DessertResolver {
	return &dessertResolver{r}
}
func (r *Resolver) Ingredient() generated.IngredientResolver {
	return &ingredientResolver{r}
}
func (r *Resolver) IngredientGroup() generated.IngredientGroupResolver {
	return &ingredientGroupResolver{r}
}
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Step() generated.StepResolver {
	return &stepResolver{r}
}
func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
