package server

import (
	"log"

	"github.com/3dw1nM0535/go-gql-server/internal/handlers"
	"github.com/gin-gonic/gin"
)

// HOST name
var HOST string

// PORT number
var PORT string

func init() {
	HOST = "localhost"
	PORT = "7777"
}

// Run web server
func Run() {
	r := gin.Default()
	// Setup routes
	r.GET("/ping", handlers.Ping())
	log.Println("Running @ http://" + HOST + ":" + PORT)
	log.Fatalln(r.Run(HOST + ":" + PORT))
}
