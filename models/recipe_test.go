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

var recipe = &m.Recipe{
	Name:         "TestRecipe",
	Description:  "TestDescription",
	Image:        "TestImage",
	RequiredTime: 10,
	Difficulty:   1,
	Servings:     1,
	Steps:        "TestSteps",
	Ingredients: []*m.Ingredient{
		{
			Name: "IngredientOne",
			Type: "TypeOne",
		},
		{
			Name: "IngredientTwo",
			Type: "TypeTwo",
		},
	},
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
	db.Exec("DELETE FROM recipes")
	db.Exec("DELETE FROM recipe_ingredients")
}

func TestRecipesTableExists(t *testing.T) {
	initializeDBRecipes()
	if (!db.HasTable(&m.Recipe{}) || !db.HasTable("recipes")) {
		t.Errorf("Recipes table wasn't created correctly")
	}
	if !db.HasTable("recipe_ingredients") {
		t.Errorf("Recipes-Ingredients relation table wasn't created correctly")
	}

	var count int
	db.Table("recipes").Count(&count)
	if count != 0 {
		t.Errorf("Recipes table isn't empty")
	}
}

func TestCreateRecipe(t *testing.T) {
	initializeDBRecipes()

	var count int
	db.Table("recipes").Count(&count)
	if count != 0 {
		t.Fatalf("Recipe table isn't empty")
	}

	err := recipe.CreateRecipe(db)
	if err != nil {
		t.Fatalf("Error creating recipe")
	}

	db.Table("recipes").Count(&count)
	if count != 1 {
		t.Errorf("Recipe wasn't created correctly")
	}
}

func TestGetRecipe(t *testing.T) {
	initializeDBRecipes()
	recipe.CreateRecipe(db)
	createdRecipe, err := m.GetRecipe(db, recipe.ID)
	if err != nil {
		t.Fatalf("Error getting recipe")
	}

	if createdRecipe.Name != "TestRecipe" {
		t.Errorf("Recipe name wasn't saved correctly")
	}
	if createdRecipe.Description != "TestDescription" {
		t.Errorf("Recipe description wasn't saved correctly")
	}
	if createdRecipe.Image != "TestImage" {
		t.Errorf("Recipe image wasn't saved correctly")
	}
	if createdRecipe.RequiredTime != 10 {
		t.Errorf("Recipe required time wasn't saved correctly")
	}
	if createdRecipe.Difficulty != 1 {
		t.Errorf("Recipe difficulty wasn't saved correctly")
	}
	if createdRecipe.Servings != 1 {
		t.Errorf("Recipe servings wasn't saved correctly")
	}
	if createdRecipe.Steps != "TestSteps" {
		t.Errorf("Recipe steps wasn't saved correctly")
	}
	if len(createdRecipe.Ingredients) != 2 {
		t.Errorf("Recipe ingredienst weren't saved correctly")
	}
}

func TestUpdateRecipe(t *testing.T) {
	initializeDBRecipes()
	recipe.CreateRecipe(db)

	recipe.Name = "NewName"
	recipe.Description = "NewDescription"
	recipe.Image = "NewImage"
	recipe.RequiredTime = 100
	recipe.Servings = 2
	recipe.Difficulty = 5
	recipe.Steps = "NewSteps"

	err := recipe.UpdateRecipe(db)
	if err != nil {
		t.Fatalf("Error updating recipe")
	}

	createdRecipe, _ := m.GetRecipe(db, recipe.ID)
	if createdRecipe.Name != "NewName" {
		t.Errorf("Recipe name wasn't saved correctly")
	}
	if createdRecipe.Description != "NewDescription" {
		t.Errorf("Recipe description wasn't saved correctly")
	}
	if createdRecipe.Image != "NewImage" {
		t.Errorf("Recipe image wasn't saved correctly")
	}
	if createdRecipe.RequiredTime != 100 {
		t.Errorf("Recipe required time wasn't saved correctly")
	}
	if createdRecipe.Difficulty != 5 {
		t.Errorf("Recipe difficulty wasn't saved correctly")
	}
	if createdRecipe.Servings != 2 {
		t.Errorf("Recipe servings weren't saved correctly")
	}
	if createdRecipe.Steps != "NewSteps" {
		t.Errorf("Recipe steps weren't saved correctly")
	}
}

func TestDeleteRecipe(t *testing.T) {
	initializeDBRecipes()
	recipe.CreateRecipe(db)

	err := m.DeleteRecipe(db, recipe.ID)
	if err != nil {
		t.Fatalf("Error deleting recipe")
	}

	var count int
	db.Table("recipes").Count(&count)
	if count != 0 {
		t.Errorf("Recipe wasn't deleted correctly")
	}
}
