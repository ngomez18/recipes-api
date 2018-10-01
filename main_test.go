package main

import (
	"bytes"
	"encoding/json"
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

func TestCreateIngredient(t *testing.T) {
	clearIngredientsTable()

	fakeName := "test ingredient"
	fakeType := "vegetable"
	payload := []byte(fmt.Sprintf(`{"name":"%s","type":"%s"}`, fakeName, fakeType))

	req, _ := http.NewRequest("POST", "/api/ingredient", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != fakeName {
		t.Errorf("Expected Ingredient name to be '%s'. Got '%v'", fakeName, m["name"])
	}

	if m["type"] != fakeType {
		t.Errorf("Expected Ingredient type to be '%s'. Got '%v'", fakeType, m["type"])
	}

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected Ingredient ID to be '1'. Got '%v'", m["id"])
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
