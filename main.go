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
	initializeIngredientRoutes(app.GetRouter())
	initializeRecipeRoutes(app.GetRouter())
	app.Run(":" + os.Getenv("PORT"))
	app.GetDB().Close()
}

func initializeIngredientRoutes(r *mux.Router) {
	r.HandleFunc("/api/ingredient", c.CreateIngredient).Methods("POST")
	r.HandleFunc("/api/ingredient/{name}", c.GetIngredient).Methods("GET")
	r.HandleFunc("/api/ingredients", c.GetIngredients).Methods("GET")
	r.HandleFunc("/api/ingredient", c.UpdateIngredient).Methods("PUT")
	r.HandleFunc("/api/ingredient/{name}", c.DeleteIngredient).Methods("DELETE")
}

func initializeRecipeRoutes(r *mux.Router) {
	r.HandleFunc("/api/recipe", c.CreateRecipe).Methods("POST")
	r.HandleFunc("/api/recipe/{id:[0-9]+}", c.GetRecipe).Methods("GET")
	r.HandleFunc("/api/recipes", c.GetRecipes).Methods("GET")
	r.HandleFunc("/api/recipe", c.UpdateRecipe).Methods("PUT")
	r.HandleFunc("/api/recipe/{id:[0-9]+}", c.DeleteRecipe).Methods("DELETE")
	r.HandleFunc("/api/recipes/name", c.GetRecipesByName).Queries("name", "{name}").Methods("GET")
}
