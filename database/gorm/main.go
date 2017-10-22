package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	// init db
	if db == nil {
		d, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			panic(err)
		}
		db = d
	}

	// migrate
	db.AutoMigrate(&product{})
}

func main() {
	http.HandleFunc("/products", handleProducts)
	http.HandleFunc("/products/one", handleSingleProduct)
	log.Fatal(http.ListenAndServe(":8001", nil))

	// graceful close the db
	defer db.Close()
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		listAllProducts(w, r)
	} else if r.Method == "POST" {
		createProduct(w, r)
	} else {
		methodNotAllwed(w, r)
	}
}

func handleSingleProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		showSingleProduct(w, r)
	} else if r.Method == "PATCH" {
		updateSingleProduct(w, r)
	} else if r.Method == "DELETE" {
		deleteSingleProduct(w, r)
	} else {
		methodNotAllwed(w, r)
	}

}

////////////////////// helpers //////////////////
