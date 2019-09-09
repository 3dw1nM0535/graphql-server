package server

import (
	log "github.com/3dw1nM0535/go-gql-server/internal/logger"

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

// Run spins the server
func Run(orm *orm.ORM) {
	log.Info("GORM_CONNECTION_DSN: ", utils.MustGet("GORM_CONNECTION_DSN"))

	endpoint := "http://" + host + ":" + port

	r := gin.Default()
	// Handlers
	// Simple keep-alive/ping handler
	r.GET("/ping", handlers.Ping())

	// GraphQL handler
	// Playground handler
	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Info("GraphQL Playground @ " + endpoint + gqlPgPath)
	}
	// Pass ORM instance to the GraphqlHandler
	r.POST(gqlPath, handlers.GraphqlHandler(orm))
	log.Info("GraphQL @ " + endpoint + gqlPath)

	// Log server status
	log.Info("Running @ " + endpoint)
	// Print out and exit(1) to the OS if server panics
	log.Fatal(r.Run(host + ":" + port))
}
