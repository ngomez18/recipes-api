package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Recipes API")

	// Initialize router
	router := mux.NewRouter()

	// Declare endpoints
	router.HandleFunc("/api/recipes", getRecipes).Methods("GET")
	router.HandleFunc("/api/recipe/{id}", getRecipe).Methods("GET")
	router.HandleFunc("/api/recipes", createRecipe).Methods("POST")
	router.HandleFunc("/api/recipe/{id}", updateRecipe).Methods("PUT")
	router.HandleFunc("/api/recipe/{id}", deleteRecipe).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getRecipes(w http.ResponseWriter, r *http.Request) {}

func getRecipe(w http.ResponseWriter, r *http.Request) {}

func createRecipe(w http.ResponseWriter, r *http.Request) {}

func updateRecipe(w http.ResponseWriter, r *http.Request) {}

func deleteRecipe(w http.ResponseWriter, r *http.Request) {}
