package main

import (
	"fmt"
	"log"
	rand "math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

const ingredientsTableCreationQuery = `CREATE TABLE IF NOT EXISTS ingredients
(
id SERIAL,
name TEXT NOT NULL,
type TEXT NOT NULL,
CONSTRAINT ingredients_pkey PRIMARY KEY (id)
)`

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize(
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_PORT"),
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	ensureIngredientsTableExists()

	code := m.Run()

	clearIngredientsTable()

	os.Exit(code)
}

func TestEmptyIngredientsTable(t *testing.T) {
	clearIngredientsTable()

	req, _ := http.NewRequest("GET", "/api/ingredients", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestFetchNonExistingIngredient(t *testing.T) {
	clearIngredientsTable()

	id := rand.Int()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/ingredient/%s", strconv.Itoa(id)), nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func ensureIngredientsTableExists() {
	if _, err := a.DB.Exec(ingredientsTableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func clearIngredientsTable() {
	a.DB.Exec("DELETE FROM ingredients")
	a.DB.Exec("ALTER SEQUENCE ingredients_id_seq RESTART WITH 1")
}
