package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	a "github.com/ngomez22/recipes-api/app"
	m "github.com/ngomez22/recipes-api/models"
	u "github.com/ngomez22/recipes-api/utils"
)

// CreateRecipe Handle a request to create a new recipe
func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var rec m.Recipe
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rec); err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := rec.CreateRecipe(a.GetDB()); err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.RespondWithJSON(w, http.StatusCreated, rec)
}

// GetRecipe Handle a request to get a specific recipe
func GetRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid recipe ID")
		return
	}

	i, err := m.GetRecipe(a.GetDB(), uint(id))
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			u.RespondWithError(w, http.StatusNotFound, "Recipe not found")
		default:
			u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	u.RespondWithJSON(w, http.StatusOK, i)
}

// GetRecipes Handle a request to get all recipes
func GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := m.GetRecipes(a.GetDB())
	if err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.RespondWithJSON(w, http.StatusOK, recipes)
}

// UpdateRecipe Handle a request to update a given recipe
func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	var rec m.Recipe
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rec); err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}

	defer r.Body.Close()

	if err := rec.UpdateRecipe(a.GetDB()); err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.RespondWithJSON(w, http.StatusOK, rec)
}

// DeleteRecipe Handle a request to delete a specific recipe
func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid recipe ID")
		return
	}

	if err := m.DeleteRecipe(a.GetDB(), uint(id)); err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// GetRecipesByName Handle a request to get recipes that match a given name
func GetRecipesByName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	recipes, err := m.GetRecipesByName(a.GetDB(), name)
	if err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.RespondWithJSON(w, http.StatusOK, recipes)
}
