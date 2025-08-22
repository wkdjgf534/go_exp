package main

import (
	"log"
	"social/internal/db"
	"social/internal/store"
)

func main() {
	addr := "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"
	conn, err := db.New(addr, 30, 30, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
