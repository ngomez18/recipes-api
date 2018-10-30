package models_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	m "github.com/ngomez22/recipes-api/models"
)

var ingredientOne = &m.Ingredient{
	Name: "IngredientOne",
	Type: "TypeOne",
}

var ingredientTwo = &m.Ingredient{
	Name: "IngredientTwo",
	Type: "TypeTwo",
}

var ingredients = [...]m.Ingredient{
	*ingredientOne,
	*ingredientTwo,
}

var recipe = &m.Recipe{
	Name:         "TestRecipe",
	Description:  "TestDescription",
	Image:        "TestImage",
	RequiredTime: 10,
	Difficulty:   1,
	Servings:     1,
	Steps:        "TestSteps",
	// Ingredients:  ingredients,
}

func initializeDBRecipes() {
	fmt.Println(os.Getenv("PORT"))
	connection := fmt.Sprintf("host=localhost port=5432 user=postgres password=admin dbname=recipes-test sslmode=disable")
	fmt.Println(connection)
	var err error
	db, err = gorm.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established with DB")
	db.AutoMigrate(&m.Recipe{}, &m.Ingredient{})
	clearRecipesTable()
}

func clearRecipesTable() {
	db.Exec("DELETE * FROM recipes")
	db.Exec("DELETE * FROM recipe_ingredients")
}

func TestRecipesTableExists(t *testing.T) {
	initializeDBRecipes()
	if (!db.HasTable(&m.Recipe{}) || !db.HasTable("recipes")) {
		t.Errorf("Recipes table wasn't created correctly")
	}
	if !db.HasTable("recipe_ingredients") {
		t.Errorf("Recipes-Ingredients relation table wasn't created correctly")
	}
}

func TestCreateRecipe(t *testing.T) {
	initializeDBRecipes()
	err := recipe.CreateIngredient(db)
	if err != nil {
		t.Fatalf("Error creating ingredient")
	}
	var count int
	db.Table("ingredients").Count(&count)
	if count != 1 {
		t.Errorf("Ingredient wasn't created correctly")
	}
}
