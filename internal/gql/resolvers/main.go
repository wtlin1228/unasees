package resolvers

import (
	"github.com/wtlin1228/go-gql-server/internal/gql"
	"github.com/wtlin1228/go-gql-server/internal/orm"
)

type Resolver struct {
	ORM *orm.ORM
}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
