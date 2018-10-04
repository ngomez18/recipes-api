package utils

import (
	"fmt"

	a "github.com/ngomez22/recipes-api/app"
	m "github.com/ngomez22/recipes-api/models"
)

// CheckTables ...
func CheckTables() {
	fmt.Println("Checking tables")
	db := a.GetDB()
	if !db.HasTable("ingredients") {
		fmt.Println("Table 'ingredients' didn't exist")
		db.CreateTable(&m.Ingredient{})
	}
}
