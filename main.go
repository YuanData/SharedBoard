package main

import (
	"database/sql"
	"log"

	"github.com/YuanData/SharedBoard/api"
	db "github.com/YuanData/SharedBoard/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:helloworld@localhost:5432/sharedboard?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can not connect to DB:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("can not start server:", err)
	}
}
