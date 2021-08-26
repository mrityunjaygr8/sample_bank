package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/mrityunjaygr8/sample_bank/api"
	db "github.com/mrityunjaygr8/sample_bank/db/sqlc"
	"github.com/mrityunjaygr8/sample_bank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := sql.Open(config.DBDriver, config.GetDBString())
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.Address)
	if err != nil {
		log.Fatal(err)
	}

}
