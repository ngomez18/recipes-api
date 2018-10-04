package models

import "github.com/jinzhu/gorm"

// Recipe model object
type Recipe struct {
	gorm.Model
	Name         string  `json:"name" gorm:"type:varchar(100);not null"`
	Description  string  `json:"description" gorm:"type:varchar(500);not null"`
	RequiredTime int     `json:"required_time" gorm:"type:integer;not null"`
	Difficulty   int     `json:"difficulty" gorm:"type:integer;not null"`
	Servings     int     `json:"servings" gorm:"type:integer;not null"`
	Steps        *string `json:"steps" gorm:"type:varchar(1000);not null"`
	Ingredients  []Ingredient
}
