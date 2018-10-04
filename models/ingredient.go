package models

import (
	"database/sql"
	"errors"

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
	if dbc := db.Create(&i); dbc.Error != nil {
		// Ingredient creation failed
		return dbc.Error
	}
	return nil
}

// GetIngredient ...
func (i *Ingredient) GetIngredient(db *sql.DB) error {
	return errors.New("Not implemented")
}

// UpdateIngredient ...
func (i *Ingredient) UpdateIngredient(db *sql.DB) error {
	return errors.New("Not implemented")
}

// DeleteIngredient ...
func (i *Ingredient) DeleteIngredient(db *sql.DB) error {
	return errors.New("Not implemented")
}

// GetIngredients ...
func GetIngredients(db *sql.DB, start, count int) ([]Ingredient, error) {
	return nil, errors.New("Not implemented")
}
