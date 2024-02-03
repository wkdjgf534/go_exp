package main

import (
	"log"

	"github.com/wkdjgf534/go_exp/courses/02_build_microservice_go/internal/database"
	"github.com/wkdjgf534/go_exp/courses/02_build_microservice_go/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize Database Client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatalf(err.Error())
	}
}
