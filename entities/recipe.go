package entities

// Recipe model object
type Recipe struct {
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	RequiredTime int         `json:"required_time"`
	Difficulty   int         `json:"difficulty"`
	Servings     int         `json:"servings"`
	Steps        *string     `json:"steps"`
	Ingredients  *Ingredient `json: "ingredients"`
}
