package handlers

import (
	"github.com/3dw1nM0535/go-gql-server/internal/gql"
	"github.com/3dw1nM0535/go-gql-server/internal/gql/resolvers"
	"github.com/99designs/gqlgen/handler"
	"github.com/3dw1nM0535/go-gql-server/internal/orm"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler defined the GQLGen GraphQL server handler
func GraphqlHandler(orm *orm.ORM) gin.HandlerFunc {
	// New ExecutableSchema and config are in the generated.go file
	c := gql.Config{
		Resolvers: &resolvers.Resolver{
			ORM: orm, // Pass ORM instance in the resolvers
		},
	}

	h := handler.GraphQL(gql.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
