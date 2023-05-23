package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"rentals-api/pkg/config"
	"rentals-api/pkg/rentals"

	_ "github.com/lib/pq"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := setupDB("config.yml")
	if err != nil {
		log.Fatalf("failed to init DB: %s", err)
		os.Exit(1)
	}
	log.Printf("successfully initialized DB.")

	err = runServer(db)
	if err != nil {
		log.Fatalf("failed to start the server: %s", err)
		os.Exit(1)
	}
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

func runServer(db *sql.DB) error {

	//initialize repo
	rentalRepo := rentals.NewRentalRepository(db)

	//initialize service
	rentalService := rentals.NewRentalService(rentalRepo)

	//register handlers
	mux := http.NewServeMux()
	log.Println("adding rental handlers")
	rentalService.AddHandlersToMux(mux)

	//run server
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("starting server at http://localhost:8080/rentals")
	err := httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
