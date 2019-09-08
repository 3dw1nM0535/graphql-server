package server

import (
	"log"

	"github.com/3dw1nM0535/go-gql-server/internal/orm"

	"github.com/gin-gonic/gin"

	"github.com/3dw1nM0535/go-gql-server/internal/handlers"
	"github.com/3dw1nM0535/go-gql-server/pkg/utils"
)

// host name
var host, gqlPath, gqlPgPath, port string
var isPgEnabled bool

func init() {
	host = utils.MustGet("GQL_SERVER_HOST")
	port = utils.MustGet("GQL_SERVER_PORT")
	gqlPath = utils.MustGet("GQL_SERVER_GRAPHQL_PATH")
	gqlPgPath = utils.MustGet("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH")
	isPgEnabled = utils.MustGetBool("GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED")
}

// Run spins web server
func Run(orm *orm.ORM) {
	log.Info("GORM_CONNECTION_DSN: ", utils.MustGet("GORM_CONNECTION_DSN"))
	endpoint := "http://" + host + ":" + port

	r := gin.Default()
	gin.ForceConsoleColor() // Force console color
	// Handlers
	// Simple keep-alive/ping handler
	r.GET("/ping", handlers.Ping())

	// GraphQL handlers
	// Playground handelrs
	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphQL Playground @ " + endpoint + gqlPgPath)
	}
	// Pass ORM instance to the GraphqlHandler
	r.POST(gqlPath, handlers.GraphqlHandler(orm))
	log.Info("GraphQL @ " + endpoint + gqlPath)

	// Log server status
	log.Println("Running @ " + endpoint)
	// Print out and exit(1) to the OS if server panics
	log.Fatalln(r.Run(host + ":" + port))
}
