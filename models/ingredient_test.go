package models_test

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	m "github.com/ngomez22/recipes-api/models"
)

var ingredient = &m.Ingredient{
	Name: "TestIngredient",
	Type: "TestType",
}
var ingredientTwo = &m.Ingredient{
	Name: "TestIngredient2",
	Type: "TestType",
}

func TestIngredientTableExists(t *testing.T) {
	InitializeDB()

	if (!db.HasTable(&m.Ingredient{}) || !db.HasTable("ingredients")) {
		t.Errorf("Table wasn't created correctly")
	}

	var count int
	db.Table("ingredients").Count(&count)
	if count != 0 {
		t.Errorf("Ingredients table isn't empty")
	}

	DropTables()
}

func TestCreateIngredient(t *testing.T) {
	InitializeDB()

	var count int
	db.Table("ingredients").Count(&count)
	if count != 0 {
		t.Fatalf("Ingredient table isn't empty")
	}

	err := ingredient.CreateIngredient(db)
	if err != nil {
		t.Fatalf("Error creating ingredient")
	}

	db.Table("ingredients").Count(&count)
	if count != 1 {
		t.Errorf("Ingredient wasn't created correctly")
	}

	DropTables()
}

func TestGetIngredient(t *testing.T) {
	InitializeDB()
	ingredient.CreateIngredient(db)

	createdIngredient, err := m.GetIngredient(db, "TestIngredient")
	if err != nil {
		t.Fatalf("Error getting ingredient")
	}

	if createdIngredient.Type != "TestType" {
		t.Errorf("Ingredient wasn't created correctly")
	}

	DropTables()
}

func TestGetIngredients(t *testing.T) {
	InitializeDB()
	ingredient.CreateIngredient(db)
	ingredientTwo.CreateIngredient(db)

	ingredients, err := m.GetIngredients(db)
	if err != nil {
		t.Fatalf("Error getting ingredients")
	}

	if len(ingredients) != 2 {
		t.Errorf("Ingredient weren't retrieved correctly")
	}

	DropTables()
}

func TestUpdateIngredient(t *testing.T) {
	InitializeDB()
	ingredient.CreateIngredient(db)

	ingredient.Type = "NewTestType"

	err := ingredient.UpdateIngredient(db)
	if err != nil {
		t.Fatalf("Error updating ingredient")
	}

	createdIngredient, _ := m.GetIngredient(db, "TestIngredient")
	if createdIngredient.Type != "NewTestType" {
		t.Errorf("Ingredient wasn't created correctly")
	}

	DropTables()
}

func TestDeleteIngredient(t *testing.T) {
	InitializeDB()
	ingredient.CreateIngredient(db)

	err := m.DeleteIngredient(db, "TestIngredient")
	if err != nil {
		t.Fatalf("Error deleting ingredient")
	}

	var count int
	db.Table("ingredients").Count(&count)
	if count != 0 {
		t.Errorf("Ingredient wasn't deleted correctly")
	}

	DropTables()
}
