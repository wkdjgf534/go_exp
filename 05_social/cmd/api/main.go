package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"social/internal/db"
	"social/internal/store"
)

const version = "0.0.1"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		log.Fatalf("error during converting string to int: %v", err)
	}

	maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		log.Fatalf("error during converting string to int: %v", err)
	}

	cfg := config{
		addr: os.Getenv("ADDR"),
		db: dbConfig{
			addr:         os.Getenv("DB_ADDR"),
			maxOpenConns: maxOpenConns,
			maxIdleConns: maxIdleConns,
			maxIdleTime:  os.Getenv("DB_MAX_IDLE_TIME"),
		},
		env: os.Getenv("ENV"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
