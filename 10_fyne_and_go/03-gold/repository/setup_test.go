package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/glebarez/go-sqlite"
)

var testRepo *SQLiteRepository

func TestMain(m *testing.M) {
	_ = os.Remove("./testdata/sql.db") // Remove temp DB for testing
	path := "./testdata/sql.db"
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Println(err)
	}

	testRepo = NewSQLiteRepository(db)
	os.Exit(m.Run())
}
