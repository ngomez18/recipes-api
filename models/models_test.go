package models_test

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	m "github.com/ngomez22/recipes-api/models"
)

var db *gorm.DB

func InitializeDB() {
	connection := fmt.Sprintf("host=localhost port=5432 user=postgres password=admin dbname=recipes-test sslmode=disable")

	var err error
	db, err = gorm.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&m.Recipe{}, &m.Ingredient{})
}

func DropTables() {
	db.Exec("DROP TABLE recipe_ingredients")
	db.Exec("DROP TABLE recipes")
	db.Exec("DROP TABLE ingredients")
}
