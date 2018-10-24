package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Blank import of DB manager package
	m "github.com/ngomez22/recipes-api/models"
)

var router *mux.Router
var db *gorm.DB

// Initialize app components
func Initialize(host, port, user, password, dbname, ssl string) {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, ssl)
	var err error
	db, err = gorm.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established with DB")
	fmt.Println(connection)

	db.LogMode(true)
	db.AutoMigrate(&m.Recipe{}, &m.Ingredient{})
	router = mux.NewRouter()
}

// Run the server
func Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}

//GetRouter ...
func GetRouter() *mux.Router {
	return router
}
