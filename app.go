package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App model
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize app components
func (a *App) Initialize(host, port, user, password, dbname string) {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run the server
func (a *App) Run(addr string) {}
