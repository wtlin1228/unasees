package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/wtlin1228/go-gql-server/internal/auth"
	gql "github.com/wtlin1228/go-gql-server/internal/gql/generated"
	"github.com/wtlin1228/go-gql-server/internal/gql/resolvers"
	"github.com/wtlin1228/go-gql-server/internal/orm"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(orm *orm.ORM) gin.HandlerFunc {

	h := auth.Middleware(
		handler.GraphQL(gql.NewExecutableSchema(resolvers.NewRootResolvers(orm))),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
