package server

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/3dw1nM0535/go-gql-server/internal/handlers"
	"github.com/3dw1nM0535/go-gql-server/pkg/utils"
)

// host name
var host string

// gqlPath
var gqlPath string

// gqlPgPath
var gqlPgPath string

// isPgEnabled
var isPgEnabled bool

// PORT number
var port string

func init() {
	host = utils.MustGet("GQL_SERVER_HOST")
	port = utils.MustGet("GQL_SERVER_PORT")
	gqlPath = utils.MustGet("GQL_SERVER_GRAPHQL_PATH")
	gqlPgPath = utils.MustGet("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH")
	isPgEnabled = utils.MustGetBool("GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED")
}

// Run web server
func Run() {
	endpoint := "http://" + host + ":" + port

	r := gin.Default()
	// Handlers
	// Simple keep-alive/ping handler
	r.GET("/ping", handlers.Ping())

	// GraphQL handlers
	// Playground handelrs
	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphLQ Playground @ " + endpoint + gqlPgPath)
	}
	r.POST(gqlPath, handlers.GraphqlHandler())

	// Log server status
	log.Println("Running @ " + endpoint)
	// Print out and exit(1) to the OS if server panics
	log.Fatalln(r.Run(host + ":" + port))
}
