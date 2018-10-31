package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	a "github.com/ngomez22/recipes-api/app"
	m "github.com/ngomez22/recipes-api/models"
	u "github.com/ngomez22/recipes-api/utils"
)

// CreateIngredient Handle a request to create a new ingredient
func CreateIngredient(w http.ResponseWriter, r *http.Request) {
	var i m.Ingredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	i.Name = strings.ToLower(i.Name)
	if err := i.CreateIngredient(a.GetDB()); err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.RespondWithJSON(w, http.StatusCreated, i)
}

// GetIngredient Handle a request to get a specific ingredient
func GetIngredient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	i, err := m.GetIngredient(a.GetDB(), string(name))
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			u.RespondWithError(w, http.StatusNotFound, "Ingredient not found")
		default:
			u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	u.RespondWithJSON(w, http.StatusOK, i)
}

// GetIngredients Handle a request to get all ingredients
func GetIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients, err := m.GetIngredients(a.GetDB())
	if err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.RespondWithJSON(w, http.StatusOK, ingredients)
}

// UpdateIngredient Handle a request to update a given ingredient
func UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	var i m.Ingredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}

	defer r.Body.Close()

	if err := i.UpdateIngredient(a.GetDB()); err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.RespondWithJSON(w, http.StatusOK, i)
}

// DeleteIngredient Handle an udpate to delete a specific ingredient
func DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if err := m.DeleteIngredient(a.GetDB(), string(name)); err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
