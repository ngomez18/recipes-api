package entities

import (
	"database/sql"
	"errors"
)

// Ingredient model object
type Ingredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// GetIngredient ...
func (i *Ingredient) GetIngredient(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM ingredients WHERE id=$1",
		i.ID).Scan(&i.Name, &i.Type)
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
	rows, err := db.Query(
		"SELECT id, name, type FROM ingredients LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ingredients := []Ingredient{}

	for rows.Next() {
		var i Ingredient
		if err := rows.Scan(&i.ID, &i.Name, &i.Type); err != nil {
			return nil, err
		}
		ingredients = append(ingredients, i)
	}
	return ingredients, nil
}
