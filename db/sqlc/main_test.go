package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	// config, err := util.LoadConfig("../..")
	// if err != nil {
	// 	log.Fatal("cannot load in th config:", err)
	// }

	conn, err := sql.Open("postgres", "postgresql://root:Maliborh521908@localhost:5432/yatzes_db?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to the database:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
