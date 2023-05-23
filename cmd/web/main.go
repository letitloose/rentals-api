package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"rentals-api/pkg/config"

	_ "github.com/lib/pq"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_, err := setupDB("config.yml")
	if err != nil {
		log.Fatalf("failed to init DB: %s", err)
		os.Exit(1)
	}
	log.Printf("successfully initialized DB.")
}

func setupDB(configFile string) (*sql.DB, error) {
	config := config.GetConfig()
	err := config.ReadConfig(configFile)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to read config file:%s", err))
	}

	connStr := config.AssembleConnectString()
	db, err := sql.Open(config.Db.Type, connStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to open DB:%s", err))
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
