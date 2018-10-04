package main

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	app "github.com/ngomez22/recipes-api/app"
	c "github.com/ngomez22/recipes-api/controllers"
)

func main() {
	fmt.Println("Starting Recipes API")

	app.Initialize(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_SSL"))
	initializeRoutes(app.GetRouter())
	app.Run(":8000")
	app.GetDB().Close()
}

func initializeRoutes(r *mux.Router) {
	r.HandleFunc("/api/ingredient", c.CreateIngredient).Methods("POST")
	r.HandleFunc("/api/ingredient/{id:[0-9]+}", c.GetIngredient).Methods("GET")
	r.HandleFunc("/api/ingredients", c.GetIngredients).Methods("GET")
	r.HandleFunc("/api/ingredient", c.UpdateIngredient).Methods("PUT")
	r.HandleFunc("/api/ingredient/{id:[0-9]+}", c.DeleteIngredient).Methods("DELETE")
}
