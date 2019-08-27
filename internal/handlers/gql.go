package handlers

import (
	"github.com/3dw1nM0535/go-gql-server/internal/gql"
	"github.com/3dw1nM0535/go-gql-server/internal/gql/resolvers"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler defined the GQLGen GraphQL server handler
func GraphqlHandler() gin.HandlerFunc {
	// New ExecutableSchema and config are in the generated.go file
	c := gql.Config{
		Resolvers: &resolvers.Resolver{},
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
