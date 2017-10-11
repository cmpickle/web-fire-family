package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)

type Product struct {
	ProductID           int     `json:"productid,omitempty"`
	ProductName         string  `json:"productname,omitempty"`
	InventoryScanningID int     `json:"inventoryscanningid,omitempty"`
	Color               string  `json:"color,omitempty"`
	Size                string  `json:"size,omitempty"`
	Price               float32 `json:"price,omitempty"`
	Dimensions          string  `json:"dimensions,omitempty"`
	SKU                 int     `json:"sku,omitempty"`
}

var Products []Product

// InitRoutes creates the web API routes and sets their event handler functions
func InitRoutes() http.Handler {
	router := mux.NewRouter()
	//This should bring a list of all the Products

	Products = append(Products, Product{ProductID: 1, ProductName: "Firefighter Wallet", InventoryScanningID: 1, Color: "Tan", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	Products = append(Products, Product{ProductID: 2, ProductName: "Firefighter Apron", InventoryScanningID: 2, Color: "Tan", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	Products = append(Products, Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", InventoryScanningID: 3, Color: "Tan", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.HandleFunc("/product", getProducts).Methods("GET")
	// This should bring back a specific Product
	router.HandleFunc("/product/{id}", getProduct).Methods("GET")
	//This creates a new product using a Json String
	router.HandleFunc("/product/create", createProduct).Methods("POST")
	//This sets the product to inactive in the database
	router.HandleFunc("/product/delete/{id}", deleteProduct).Methods("DELETE")
	//This allows us to set the quantity value of a product.
	router.HandleFunc("/inventory/update/{id}/{quantity}", updateInventory).Methods("POST")
	return router
}

// Returns all of the products stored in the database in JSON format
func getProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Products)
}

// Returns a specific product from the database in JSON format
func getProduct(w http.ResponseWriter, r *http.Request) {

}

// Creates a Product object from the passed in JSON Product and stores it in the database
func createProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ProductID, _ = strconv.Atoi(params["productid"])
	products = append(products, product)
	json.NewEncoder(w).Encode(products)

}

// Deletes the specified product from the database
func deleteProduct(w http.ResponseWriter, r *http.Request) {

}

// Updates the inventory value for the inventory item
func updateInventory(w http.ResponseWriter, r *http.Request) {

}
