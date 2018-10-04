package models

import (
	"github.com/jinzhu/gorm"
)

// Ingredient model object
type Ingredient struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(100);not null"`
	Type string `json:"type" gorm:"type:varchar(100);not null"`
}

// CreateIngredient ...
func (i *Ingredient) CreateIngredient(db *gorm.DB) error {
	if response := db.Create(&i); response.Error != nil {
		// Ingredient creation failed
		return response.Error
	}
	return nil
}

// GetIngredient ...
func GetIngredient(db *gorm.DB, id uint) (*Ingredient, error) {
	ingredient := &Ingredient{}
	response := db.First(ingredient, id)
	return ingredient, response.Error
}

// GetIngredients ...
func GetIngredients(db *gorm.DB, start, count int) ([]Ingredient, error) {
	ingredients := make([]Ingredient, 0)
	response := db.Limit(count).Offset(start).Find(&ingredients)
	return ingredients, response.Error
}

// UpdateIngredient ...
func (i *Ingredient) UpdateIngredient(db *gorm.DB) error {
	if response := db.Save(&i); response.Error != nil {
		// Ingredient update failed
		return response.Error
	}
	return nil
}

// DeleteIngredient ...
func DeleteIngredient(db *gorm.DB, id uint) error {
	if response := db.Where("id = ?", id).Delete(&Ingredient{}); response.Error != nil {
		// Ingredient delete failed
		return response.Error
	}
	return nil
}
