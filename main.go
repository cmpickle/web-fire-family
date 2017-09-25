package web_fire_family


import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
 router := mux.NewRouter()
 //This should bring a list of all the Products
 router.HandleFunc("/product", GetProducts).Methods("GET")
 // This should bring back a specific Product
 router.HandleFunc("/product/{id}", GetProduct).Methods("GET")
 //This creates a new product using a Json String
 router.HandleFunc("/product/Create",CreateProduct).Methods("POST")
 //This sets the product to inactive in the database
 router.HandleFunc("/product/delete/{id}", DeleteProduct).Methods("DELETE")
 //This allows us to set the quantity value of a product.
 router.HandleFunc("/inventory/update/{id}/{quantity}",UpdateInventory).Methods("POST")
 log.Fatal(http.ListenAndServe(":8000", router))
}