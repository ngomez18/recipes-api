package entities

import (
	"database/sql"
	"errors"
)

// Ingredient model object
type Ingredient struct {
	Name string `json:"name"`
	Type string `json:"type"`
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

// CreateIngredient ...
func (i *Ingredient) CreateIngredient(db *sql.DB) error {
	return errors.New("Not implemented")
}

// GetIngredients ...
func GetIngredients(db *sql.DB, start, count int) ([]Ingredient, error) {
	return nil, errors.New("Not implemented")
}
