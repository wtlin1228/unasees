package resolvers

import (
	"github.com/wtlin1228/go-gql-server/internal/gql/generated"
	"github.com/wtlin1228/go-gql-server/internal/orm"
)

type Resolver struct {
	ORM *orm.ORM
}

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

type queryResolver struct{ *Resolver }
