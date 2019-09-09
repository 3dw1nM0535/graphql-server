package main

import (
	log "github.com/3dw1nM0535/go-gql-server/internal/logger"
	"github.com/3dw1nM0535/go-gql-server/internal/orm"
	"github.com/3dw1nM0535/go-gql-server/pkg/server"
)

func main() {
	// Create a new ORM instance to pass it to our server
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}
	// Send ORM instance to our server
	server.Run(orm)
}
