package resolvers

import (
	"context"

	"github.com/wtlin1228/go-gql-server/internal/gql/generated"
	gqlmodels "github.com/wtlin1228/go-gql-server/internal/gql/models"
	"github.com/wtlin1228/go-gql-server/internal/orm/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

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

type categoryResolver struct{ *Resolver }

func (r *categoryResolver) ID(ctx context.Context, obj *models.Category) (string, error) {
	panic("not implemented")
}
func (r *categoryResolver) DessertList(ctx context.Context, obj *models.Category) ([]*models.Dessert, error) {
	panic("not implemented")
}

type dessertResolver struct{ *Resolver }

func (r *dessertResolver) ID(ctx context.Context, obj *models.Dessert) (string, error) {
	panic("not implemented")
}
func (r *dessertResolver) StepList(ctx context.Context, obj *models.Dessert) ([]*models.Step, error) {
	panic("not implemented")
}
func (r *dessertResolver) IngredientGroupList(ctx context.Context, obj *models.Dessert) ([]*models.IngredientGroup, error) {
	panic("not implemented")
}
func (r *dessertResolver) Category(ctx context.Context, obj *models.Dessert) (*models.Category, error) {
	panic("not implemented")
}

type ingredientResolver struct{ *Resolver }

func (r *ingredientResolver) ID(ctx context.Context, obj *models.Ingredient) (string, error) {
	panic("not implemented")
}
func (r *ingredientResolver) IngredientGroup(ctx context.Context, obj *models.Ingredient) (*models.IngredientGroup, error) {
	panic("not implemented")
}

type ingredientGroupResolver struct{ *Resolver }

func (r *ingredientGroupResolver) ID(ctx context.Context, obj *models.IngredientGroup) (string, error) {
	panic("not implemented")
}
func (r *ingredientGroupResolver) IngredientList(ctx context.Context, obj *models.IngredientGroup) ([]*models.Ingredient, error) {
	panic("not implemented")
}
func (r *ingredientGroupResolver) Dessert(ctx context.Context, obj *models.IngredientGroup) (*models.Dessert, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateCategory(ctx context.Context, input *gqlmodels.CategoryInput) (*models.Category, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateCategory(ctx context.Context, id string, input *gqlmodels.CategoryInput) (*models.Category, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateDessert(ctx context.Context, input *gqlmodels.DessertInput) (*models.Dessert, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateDessert(ctx context.Context, id string, input *gqlmodels.DessertInput) (*models.Dessert, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteDessert(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateIngredientGroup(ctx context.Context, input *gqlmodels.IngredientGroupInput) (*models.IngredientGroup, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateIngredientGroup(ctx context.Context, id string, input *gqlmodels.IngredientGroupInput) (*models.IngredientGroup, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteIngredientGroup(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateIngredient(ctx context.Context, input *gqlmodels.IngredientInput) (*models.Ingredient, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateIngredient(ctx context.Context, id string, input *gqlmodels.IngredientInput) (*models.Ingredient, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteIngredient(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateStep(ctx context.Context, input *gqlmodels.StepInput) (*models.Step, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateStep(ctx context.Context, id string, input *gqlmodels.StepInput) (*models.Step, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteStep(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) Login(ctx context.Context, input *gqlmodels.UserInput) (string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Categories(ctx context.Context) ([]*models.Category, error) {
	panic("not implemented")
}
func (r *queryResolver) Category(ctx context.Context, id string) (*models.Category, error) {
	panic("not implemented")
}
func (r *queryResolver) Desserts(ctx context.Context) ([]*models.Dessert, error) {
	panic("not implemented")
}
func (r *queryResolver) Dessert(ctx context.Context, id string) (*models.Dessert, error) {
	panic("not implemented")
}
func (r *queryResolver) IngredientGroups(ctx context.Context) ([]*models.IngredientGroup, error) {
	panic("not implemented")
}
func (r *queryResolver) IngredientGroup(ctx context.Context, id string) (*models.IngredientGroup, error) {
	panic("not implemented")
}
func (r *queryResolver) Ingredients(ctx context.Context) ([]*models.Ingredient, error) {
	panic("not implemented")
}
func (r *queryResolver) Ingredient(ctx context.Context, id string) (*models.Ingredient, error) {
	panic("not implemented")
}
func (r *queryResolver) Steps(ctx context.Context) ([]*models.Step, error) {
	panic("not implemented")
}
func (r *queryResolver) Step(ctx context.Context, id string) (*models.Step, error) {
	panic("not implemented")
}

type stepResolver struct{ *Resolver }

func (r *stepResolver) ID(ctx context.Context, obj *models.Step) (string, error) {
	panic("not implemented")
}
func (r *stepResolver) Dessert(ctx context.Context, obj *models.Step) (*models.Dessert, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	panic("not implemented")
}
