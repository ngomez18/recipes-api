package models

import "github.com/jinzhu/gorm"

// Recipe model object
type Recipe struct {
	gorm.Model
	Name         string        `json:"name" gorm:"type:varchar(100);not null"`
	Description  string        `json:"description" gorm:"type:varchar(500);not null"`
	RequiredTime int           `json:"required_time" gorm:"type:integer;not null"`
	Difficulty   int           `json:"difficulty" gorm:"type:integer;not null"`
	Servings     int           `json:"servings" gorm:"type:integer;not null"`
	Steps        *string       `json:"steps" gorm:"type:varchar(1000);not null"`
	Ingredients  []*Ingredient `gorm:"many2many:recipe_ingredients"`
}

// CreateRecipe ...
func (r *Recipe) CreateRecipe(db *gorm.DB) error {
	if response := db.Create(&r).Save(&r); response.Error != nil {
		// Recipe creation failed
		return response.Error
	}
	return nil
}

// GetRecipe ...
func GetRecipe(db *gorm.DB, id uint) (*Recipe, error) {
	recipe := &Recipe{}
	response := db.Preload("Ingredients").First(recipe, id)
	return recipe, response.Error
}

// GetRecipes ...
func GetRecipes(db *gorm.DB, start, count int) ([]Recipe, error) {
	recipes := make([]Recipe, 0)
	response := db.Limit(count).Offset(start).Preload("Ingredients").Find(&recipes)
	return recipes, response.Error
}

// UpdateRecipe ...
func (r *Recipe) UpdateRecipe(db *gorm.DB) error {
	if response := db.Save(&r); response.Error != nil {
		// Recipe update failed
		return response.Error
	}
	return nil
}

// DeleteRecipe ...
func DeleteRecipe(db *gorm.DB, id uint) error {
	if response := db.Where("id = ?", id).Delete(&Recipe{}); response.Error != nil {
		// Recipe delete failed
		return response.Error
	}
	return nil
}
