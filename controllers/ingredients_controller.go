package controllers

import (
	"encoding/json"
	"net/http"

	a "github.com/ngomez22/recipes-api/app"
	m "github.com/ngomez22/recipes-api/models"
	u "github.com/ngomez22/recipes-api/utils"
)

// CreateIngredient ...
func CreateIngredient(w http.ResponseWriter, r *http.Request) {
	var i m.Ingredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := i.CreateIngredient(a.GetDB()); err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.RespondWithJSON(w, http.StatusCreated, i)
}
