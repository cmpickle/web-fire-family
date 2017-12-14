package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cmpickle/web-fire-family/app"
	"github.com/cmpickle/web-fire-family/models"
)

var db *sql.DB
var settings app.Dbdriver
var dbConnection string

// InitRoutes creates the web API routes and sets their event handler functions
func InitRoutes(env models.Env) http.Handler {
	router := mux.NewRouter()

	db = env.Db

	// Bootstrapping the setting

	router.HandleFunc("/product", getProducts).Methods("GET")
	// This should bring back a specific Product.
	router.HandleFunc("/product/{sku}", getProductBySKU).Methods("GET")
	// This should bring back a specific Product.
	//router.HandleFunc("/productbyid/{id}", getProduct).Methods("GET")
	//This creates a new product using a Json String.
	router.HandleFunc("/product/create", createProduct).Methods("POST")
	//This updates a product using a Json String.
	router.HandleFunc("/product/update/{sku}", updateProductBySKU).Methods("POST")
	//This sets the product to inactive in the database.
	router.HandleFunc("/product/delete/{sku}", deleteProductBySKU).Methods("POST")
	//This gets the inventory values.
	router.HandleFunc("/inventories", getInventories).Methods("GET")
	//This gets the inventory value.
	router.HandleFunc("/inventory/{sku}", getInventoryBySKU).Methods("GET")
	// This should bring back a specific Inventory.
	//router.HandleFunc("/inventorybyid/{id}", getInventory).Methods("GET")
	//This allows the quantity value of a product to be set.
	router.HandleFunc("/inventory/update/{sku}/{quantity}", updateInventoryBySKU).Methods("POST")
	//This allows for incrementation of a product's inventory.
	router.HandleFunc("/inventory/increment/{sku}", incrementInventoryBySKU).Methods("POST")
	//This allows for decrementation of a product's inventory.
	router.HandleFunc("/inventory/decrement/{sku}", decrementInventoryBySKU).Methods("POST")

	return router
}
