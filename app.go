package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/ngomez22/recipes-api/entities"
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
	fmt.Println(connection)
	var err error
	a.DB, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run the server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) createIngredient(w http.ResponseWriter, r *http.Request) {
	var i entities.Ingredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := i.CreateIngredient(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, i)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/ingredient", a.createIngredient).Methods("POST")
}

// Responde with a JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Respond with an error JSON
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
