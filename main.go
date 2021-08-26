package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/mrityunjaygr8/sample_bank/api"
	db "github.com/mrityunjaygr8/sample_bank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/sample_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start("localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

}
