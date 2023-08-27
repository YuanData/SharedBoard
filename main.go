package main

import (
	"database/sql"
	"log"

	"github.com/YuanData/SharedBoard/api"
	"github.com/YuanData/SharedBoard/cfg"
	db "github.com/YuanData/SharedBoard/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	config, err := cfg.LoadConfig("./cfg")
	if err != nil {
		log.Fatal("can not load configuration:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to DB:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("can not start server:", err)
	}
}
