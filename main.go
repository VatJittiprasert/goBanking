package main

import (
	"database/sql"
	"log"

	"github.com/VatJittiprasert/goBanking/api"
	db "github.com/VatJittiprasert/goBanking/db/sqlc"
	"github.com/VatJittiprasert/goBanking/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot laod config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
