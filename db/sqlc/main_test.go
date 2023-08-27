package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/YuanData/SharedBoard/cfg"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := cfg.LoadConfig("../../cfg")
	if err != nil {
		log.Fatal("can not load configuration:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to DB:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
