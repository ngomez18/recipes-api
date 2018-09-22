package entities

// Ingredient model object
type Ingredient struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Units    string `json:"units"`
}
