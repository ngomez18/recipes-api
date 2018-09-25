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

func (p *Ingredient) getIngredient(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Ingredient) updateIngredient(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Ingredient) deleteIngredient(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Ingredient) createIngredient(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getIngredients(db *sql.DB, start, count int) ([]Ingredient, error) {
	return nil, errors.New("Not implemented")
}
