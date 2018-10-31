package models_test

import (
	"testing"

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

func TestRecipesTableExists(t *testing.T) {
	InitializeDB()

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

	DropTables()
}

func TestCreateRecipe(t *testing.T) {
	InitializeDB()

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

	DropTables()
}

func TestGetRecipe(t *testing.T) {
	InitializeDB()
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

	DropTables()
}

func TestGetRecipes(t *testing.T) {
	InitializeDB()
	recipe.CreateRecipe(db)

	recipes, err := m.GetRecipes(db)

	if err != nil {
		t.Fatalf("Error fetching recipes")
	}

	if len(recipes) != 1 {
		t.Errorf("Recipes weren't retrieved correctly")
	}

	DropTables()
}

func TestUpdateRecipe(t *testing.T) {
	InitializeDB()
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

	updatedRecipe, _ := m.GetRecipe(db, recipe.ID)
	if updatedRecipe.Name != "NewName" {
		t.Errorf("Recipe name wasn't saved correctly")
	}
	if updatedRecipe.Description != "NewDescription" {
		t.Errorf("Recipe description wasn't saved correctly")
	}
	if updatedRecipe.Image != "NewImage" {
		t.Errorf("Recipe image wasn't saved correctly")
	}
	if updatedRecipe.RequiredTime != 100 {
		t.Errorf("Recipe required time wasn't saved correctly")
	}
	if updatedRecipe.Difficulty != 5 {
		t.Errorf("Recipe difficulty wasn't saved correctly")
	}
	if updatedRecipe.Servings != 2 {
		t.Errorf("Recipe servings weren't saved correctly")
	}
	if updatedRecipe.Steps != "NewSteps" {
		t.Errorf("Recipe steps weren't saved correctly")
	}

	DropTables()
}

func TestDeleteRecipe(t *testing.T) {
	InitializeDB()
	recipe.CreateRecipe(db)

	err := m.DeleteRecipe(db, recipe.ID)
	if err != nil {
		t.Fatalf("Error deleting recipe")
	}

	recipes, _ := m.GetRecipes(db)

	if len(recipes) != 0 {
		t.Errorf("Recipe wasn't deleted correctly")
	}

	DropTables()
}

func TestUpdateRecipeIngredients(t *testing.T) {
	InitializeDB()
	recipe.CreateRecipe(db)

	newIngredients := []*m.Ingredient{
		{
			Name: "NewIngredientOne",
			Type: "TypeOne",
		},
		{
			Name: "NewIngredientTwo",
			Type: "TypeTwo",
		},
		{
			Name: "IngredientThree",
			Type: "TypeThree",
		},
	}
	recipe.Ingredients = newIngredients

	err := recipe.UpdateRecipe(db)
	if err != nil {
		t.Fatalf("Error updating recipe")
	}

	createdRecipe, _ := m.GetRecipe(db, recipe.ID)
	if len(createdRecipe.Ingredients) != len(newIngredients) {
		t.Errorf("Recipe's ingredients weren't udpated correctly")
	}

	for i := range createdRecipe.Ingredients {
		if createdRecipe.Ingredients[i].Name != newIngredients[i].Name {
			t.Errorf("Recipe's ingredients weren't udpated correctly")
		}
		if createdRecipe.Ingredients[i].Type != newIngredients[i].Type {
			t.Errorf("Recipe's ingredients weren't udpated correctly")
		}
	}

	DropTables()
}
