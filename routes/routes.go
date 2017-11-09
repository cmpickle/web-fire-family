package routes

import (
	"database/sql"
	//"log"
	"net/http"

	"github.com/gorilla/mux"

	//"github.com/Xero67/web-fire-family/app"
	//"github.com/Xero67/web-fire-family/models"
	"../models"
	"../app"


	_ "github.com/go-sql-driver/mysql"
	//"image/color"
)

var Products []models.Product

var db *sql.DB
var settings app.Dbdriver
var dbConnection string

// InitRoutes creates the web API routes and sets their event handler functions
func InitRoutes(env models.Env) http.Handler {
	router := mux.NewRouter()

	db = env.Db

	//This should bring a list of all the Products
	Products = append(Products, models.Product{ProductID: 1, ProductName: "Firefighter Wallet", NotificationQuantity: 10, Color: "Tan", TrimColor: "Black", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	Products = append(Products, models.Product{ProductID: 2, ProductName: "Firefighter Apron", NotificationQuantity: 20, Color: "Tan", TrimColor: "Black", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	Products = append(Products, models.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", NotificationQuantity: 13, Color: "Tan", TrimColor: "Black", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	// Bootstrapping the setting

	router.HandleFunc("/product", getProducts).Methods("GET")
	// This should bring back a specific Product
	router.HandleFunc("/product/{id}", getProduct).Methods("GET")
	//This creates a new product using a Json String
	router.HandleFunc("/product/create", createProduct).Methods("POST")
	//This updates a product using a Json String
	router.HandleFunc("/product/update/{id}", updateProduct).Methods("PUT")
	//This sets the product to inactive in the database
	router.HandleFunc("/product/delete/{id}", deleteProduct).Methods("DELETE")
	//This gets the inventory values
	router.HandleFunc("/inventories", getInventories).Methods("GET")
	//This gets the inventory value
	router.HandleFunc("/inventory/{id}", getInventory).Methods("GET")
	//This allows us to set the quantity value of a product.
	router.HandleFunc("/inventory/update/{id}/{quantity}", updateInventory).Methods("PUT")

	return router
}
