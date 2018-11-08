package models

import "github.com/jinzhu/gorm"

// Recipe model object
type Recipe struct {
	gorm.Model
	Name         string        `json:"name" gorm:"type:varchar(100);not null"`
	Description  string        `json:"description" gorm:"type:varchar(500);not null"`
	Image        string        `json:"image" gorm:"type:varchar(2000);not null"`
	RequiredTime int           `json:"requiredTime" gorm:"type:integer;not null"`
	Difficulty   int           `json:"difficulty" gorm:"type:integer;not null"`
	Servings     int           `json:"servings" gorm:"type:integer;not null"`
	Steps        string        `json:"steps" gorm:"type:varchar(1000);not null"`
	Ingredients  []*Ingredient `json:"ingredients" gorm:"many2many:recipe_ingredients"`
}

// CreateRecipe Save the given recipe to the database
func (r *Recipe) CreateRecipe(db *gorm.DB) error {
	if response := db.Create(&r).Save(&r); response.Error != nil {
		// Recipe creation failed
		return response.Error
	}
	return nil
}

// GetRecipe Get the recipe with the given ID from the database
func GetRecipe(db *gorm.DB, id uint) (*Recipe, error) {
	recipe := &Recipe{}
	response := db.Preload("Ingredients").First(recipe, id)
	return recipe, response.Error
}

// GetRecipes Get all recipes from the database
func GetRecipes(db *gorm.DB) ([]Recipe, error) {
	recipes := make([]Recipe, 0)
	response := db.Preload("Ingredients").Find(&recipes)
	return recipes, response.Error
}

// UpdateRecipe Update the given recipe in the database
func (r *Recipe) UpdateRecipe(db *gorm.DB) error {
	if response := db.Save(&r); response.Error != nil {
		// Recipe update failed
		return response.Error
	}
	if response := db.Model(r).Association("Ingredients").Replace(r.Ingredients); response.Error != nil {
		// Recipe's ingredients update failed
		return response.Error
	}
	return nil
}

// DeleteRecipe Delete the recipe with the given ID from the database
func DeleteRecipe(db *gorm.DB, id uint) error {
	if response := db.Where("id = ?", id).Delete(&Recipe{}); response.Error != nil {
		// Recipe delete failed
		return response.Error
	}
	return nil
}

// GetRecipesByName Get recipes that match a given name
func GetRecipesByName(db *gorm.DB, name string) ([]Recipe, error) {
	recipes := make([]Recipe, 0)
	response := db.Preload("Ingredients").Where("name ILIKE ?", name+"%").Find(&recipes)
	return recipes, response.Error
}
