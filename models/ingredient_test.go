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

var db *gorm.DB
var ingredient = &m.Ingredient{
	Name: "TestIngredient",
	Type: "TestType",
}

func initializeDB() {
	fmt.Println(os.Getenv("PORT"))
	connection := fmt.Sprintf("host=localhost port=5432 user=postgres password=admin dbname=recipes-test sslmode=disable")
	fmt.Println(connection)
	var err error
	db, err = gorm.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established with DB")
	db.AutoMigrate(&m.Ingredient{})
	clearIngredientsTable()
}

func clearIngredientsTable() {
	db.Exec("DELETE * FROM ingredients")
}

func TestIngredientTableExists(t *testing.T) {
	initializeDB()
	if (!db.HasTable(&m.Ingredient{}) || !db.HasTable("ingredients")) {
		t.Errorf("Table wasn't created correctly")
	}
}

func TestCreateIngredient(t *testing.T) {
	initializeDB()
	ingredient.CreateIngredient(db)
	var count int
	db.Table("ingredients").Count(&count)
	if count != 1 {
		t.Errorf("Ingredient wasn't created correctly")
	}
}

func TestGetIngredient(t *testing.T) {
	initializeDB()
	ingredient.CreateIngredient(db)
	createdIngredient, err := m.GetIngredient(db, "TestIngredient")
	if err != nil {
		t.Fatalf("Error getting ingredient")
	}
	if createdIngredient.Type != "TestType" {
		t.Errorf("Ingredient wasn't created correctly")
	}
}
