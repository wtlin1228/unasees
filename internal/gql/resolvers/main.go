package resolvers

import (
	"github.com/wtlin1228/go-gql-server/internal/gql/generated"
	"github.com/wtlin1228/go-gql-server/internal/orm"
)

type Resolver struct {
	ORM *orm.ORM
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

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
