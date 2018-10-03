package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	e "github.com/ngomez22/recipes-api/entities"
	u "github.com/ngomez22/recipes-api/utils"
)

// App model
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize app components
func (a *App) Initialize(host, port, user, password, dbname, ssl string) {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, ssl)
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
	var i e.Ingredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := i.CreateIngredient(a.DB); err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.RespondWithJSON(w, http.StatusCreated, i)
}

func (a *App) getIngredients(w http.ResponseWriter, r *http.Request) {
	u.RespondWithJSON(w, http.StatusCreated, e.Ingredient{Name: "Ing1", Type: "Food"})
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/ingredients", a.getIngredients).Methods("GET")
	a.Router.HandleFunc("/ingredient", a.createIngredient).Methods("POST")
}
