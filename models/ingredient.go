package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Ingredient model object
type Ingredient struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name" gorm:"primary_key;type:varchar(100);not null"`
	Type      string `json:"type" gorm:"type:varchar(100);not null"`
}

// CreateIngredient Save the given ingredient to the database
func (i *Ingredient) CreateIngredient(db *gorm.DB) error {
	if response := db.Create(&i); response.Error != nil {
		// Ingredient creation failed
		return response.Error
	}
	return nil
}

// GetIngredient Get the ingredient with the given name from the database
func GetIngredient(db *gorm.DB, name string) (*Ingredient, error) {
	ingredient := &Ingredient{}
	response := db.Where("name = ?", name).First(ingredient)
	return ingredient, response.Error
}

// GetIngredients Get all the ingredients from the database
func GetIngredients(db *gorm.DB, start, count int) ([]Ingredient, error) {
	ingredients := make([]Ingredient, 0)
	response := db.Limit(count).Offset(start).Find(&ingredients)
	return ingredients, response.Error
}

// UpdateIngredient Update the given ingredient in the database
func (i *Ingredient) UpdateIngredient(db *gorm.DB) error {
	if response := db.Save(&i); response.Error != nil {
		// Ingredient update failed
		return response.Error
	}
	return nil
}

// DeleteIngredient Delete the ingredient with the given name from the database
func DeleteIngredient(db *gorm.DB, name string) error {
	if response := db.Where("name = ?", name).Delete(&Ingredient{}); response.Error != nil {
		// Ingredient delete failed
		return response.Error
	}
	return nil
}
