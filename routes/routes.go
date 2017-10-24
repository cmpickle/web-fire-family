package routes

import (
	"encoding/json"
	//"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"database/sql"
	"fmt"

	//Cannot get this import to work
	"../models"

	_ "github.com/go-sql-driver/mysql"

	//"image/color"
)
//Doesn't match our product table as of 10/20
//type Product struct {
//	ProductID           int     `json:"productid,omitempty"`
//	ProductName         string  `json:"productname,omitempty"`
//	InventoryScanningID int     `json:"inventoryscanningid,omitempty"`
//	Color               string  `json:"color,omitempty"`
//	Size                string  `json:"size,omitempty"`
//	Price               float32 `json:"price,omitempty"`
//	Dimensions          string  `json:"dimensions,omitempty"`
//	SKU                 int     `json:"sku,omitempty"`
//}

//Matches our product table
//type Productt struct {
//	ProductID           	int     `json:"productid,omitempty"`
//	ProductName         	string  `json:"productname,omitempty"`
//	NotificationQuantity	int 	`json:"notificationquantity, omitempty"`
//	Color               	string  `json:"color,omitempty"`
//	TrimColor				string	`json:"trimcolor,omitempty"`
//	Size                	string  `json:"size,omitempty"`
//	Price               	float32 `json:"price,omitempty"`
//	Dimensions          	string  `json:"dimensions,omitempty"`
//	SKU                 	int     `json:"sku,omitempty"`
//	Deleted                 int     `json:"deleted,omitempty"`
//}



var Products []models.Product
var db *sql.DB

// InitRoutes creates the web API routes and sets their event handler functions
func InitRoutes() http.Handler {
	router := mux.NewRouter()

	//Trying DB things here
	var err error
	//db, err = sql.Open("mysql", "fireadmin:FireFamily@1@tcp(165.227.17.104:3306)/Fire_Family")
	models.NewDB("fireadmin:FireFamily@1@tcp(165.227.17.104:3306)/Fire_Family")
	db = models.Db
	if err != nil {
		//error handling here
		fmt.Println("Conn")
		fmt.Println(err)
	}
	if err = db.Ping(); err != nil {
		//error handling here
		fmt.Println("Ping")
		fmt.Println(err)
	}
	//This should bring a list of all the Products

	Products = append(Products, models.Product{ProductID: 1, ProductName: "Firefighter Wallet", NotificationQuantity: 1, Color: "Tan", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	Products = append(Products, models.Product{ProductID: 2, ProductName: "Firefighter Apron", NotificationQuantity: 2, Color: "Tan", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	Products = append(Products, models.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", NotificationQuantity: 3, Color: "Tan", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.HandleFunc("/product", getProducts).Methods("GET")
	// This should bring back a specific Product
	router.HandleFunc("/product/{id}", getProduct).Methods("GET")
	//This creates a new product using a Json String
	router.HandleFunc("/product/create", createProduct).Methods("POST")
	//This updates a product using a Json String
	router.HandleFunc("/product/update/{id}", updateProduct).Methods("PUT")
	//This sets the product to inactive in the database
	router.HandleFunc("/product/delete/{id}", deleteProduct).Methods("DELETE")
	//This allows us to set the quantity value of a product.
	router.HandleFunc("/inventory/update/{id}/{quantity}", updateInventory).Methods("PUT")


	return router
}

// Returns all of the products stored in the database in JSON format
func getProducts(w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode(Products)
	rows, err := db.Query("SELECT * FROM Product")
	if err != nil {
		//Error handling
		fmt.Println("1")
		fmt.Println(err)
	}
	defer rows.Close()
	prods := make([]*models.Product, 0)
	for rows.Next() {
		p := new(models.Product)
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.NotificationQuantity, &p.Color, &p.TrimColor, &p.Size, &p.Price, &p.Dimensions, &p.SKU, &p.Deleted)
		if err != nil {
			//More error handling
			fmt.Println("2")
			fmt.Println(err)
		}
		if p.Deleted == 0 {
			prods = append(prods, p)
		}
	}
	if err = rows.Err(); err != nil {
		//Error handling
		fmt.Println("3")
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(prods)

}

// Returns a specific product from the database in JSON format
func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	found := -1

	//new stuff can easily change to work off of SKU
	productID, err := strconv.Atoi(id)
	if  productID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product ID."))
		return
	}
	rows, err := db.Query("SELECT * FROM Product WHERE ProductID = ?", id)
	if err != nil {
		//Error handling
		fmt.Println("1")
		fmt.Println(err)
	}
	defer rows.Close()
	prods := make([]*models.Product, 0)
	for rows.Next() {

		p := new(models.Product)
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.NotificationQuantity, &p.Color, &p.TrimColor, &p.Size, &p.Price, &p.Dimensions, &p.SKU, &p.Deleted)
		if err != nil {
			//More error handling
			fmt.Println("2")
			fmt.Println(err)
		}
		if p.Deleted == 0 {
			prods = append(prods, p)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("404 - Product not found"))
			return
		}
		found = productID

	}
	if err = rows.Err(); err != nil {
		//Error handling
		fmt.Println("3")
		fmt.Println(err)
	}
	//end new stuff

	/*for i, value := range Products {
		productID, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
			json.NewEncoder(w).Encode("")
		}

		if productID < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Invalid product ID."))
			return
		}

		if value.ProductID == productID {
			index = i
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Products[index])
			return
		}
	}*/

	//STILL NEED THIS FOR IF ITS NOT FOUND
	if found == -1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("404 - Product not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prods)
}

// Smart thing to do would be to check the DB for the item already being created first
// As well as if that item was already deleted then just toggle it instead
// Not needed tho, def an extra thing
// Creates a Product object from the passed in JSON Product and stores it in the database
func createProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	fmt.Println(product)

	//new stuff
	//probs want to check each required column
	if product.ProductName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product, please include a name, notification quantity, color, trim color, size, price, dimensions, and SKU"))
		return
	}
	stmnt, err := db.Prepare("INSERT INTO Product (ProductName, NotificationQuantity, Color, TrimColor, Size, Price, Dimensions, SKU) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product, please include a name, notification quantity, color, trim color, size, price, dimensions, and SKU"))
		return
	}
	res, err := stmnt.Exec(product.ProductName, product.NotificationQuantity, product.Color, product.TrimColor, product.Size, product.Price, product.Dimensions, product.SKU)
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Insert failed"))
		return
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("3")
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Insert failed"))
		return
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		fmt.Println("4")
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Insert failed"))
		return
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	lstId := strconv.Itoa(int(lastId))
	//Not sure what we want to return when sucess?
	w.Write([]byte("Success! The index is " + lstId))

	//end new stuff
	//Products = append(Products, product)
}

// Deletes the specified product from the database
// If the route logic were seperate from the DB logic, we could just call a getproductbyID method that is used
// by both
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	/*params := mux.Vars(r)
	id := params["id"]

	for i, value := range Products {
		productID, err := strconv.Atoi(id)
		if err != nil || productID < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Invalid product ID."))
			return
		}

		if value.ProductID == productID {
			Products[i] = Products[len(Products)-1]
			Products = Products[:len(Products)-1]
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 - Invalid product ID."))*/

	//new version
	params := mux.Vars(r)
	id := params["id"]
	found := -1

	productID, err := strconv.Atoi(id)
	if  productID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product ID."))
		return
	}
	rows, err := db.Query("SELECT * FROM Product WHERE ProductID = ?", id)
	if err != nil {
		//Error handling
		fmt.Println("1")
		fmt.Println(err)
	}
	defer rows.Close()
	prods := make([]*models.Product, 0)
	for rows.Next() {

		p := new(models.Product)
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.NotificationQuantity, &p.Color, &p.TrimColor, &p.Size, &p.Price, &p.Dimensions, &p.SKU, &p.Deleted)
		if err != nil {
			//More error handling
			fmt.Println("2")
			fmt.Println(err)
		}
		if p.Deleted == 0 {
			prods = append(prods, p)
		}
		found = productID

	}
	if err = rows.Err(); err != nil {
		//Error handling
		fmt.Println("3")
		fmt.Println(err)
	}

	if found == -1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product ID."))
		return
	} else { //All deletion logic goes here because it confirms the find
		prods[0].Deleted = 1;
		stmt, err := db.Prepare("UPDATE Product SET Deleted = 1 WHERE ProductID = ?")
		if err != nil {
			fmt.Println(err)
		}
		_, errr := stmt.Exec(prods[0].ProductID)
		if errr != nil {
			fmt.Println(err)
		}

	}
	w.WriteHeader(http.StatusOK)
}

// Updates the product
func updateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	params := mux.Vars(r)
	id := params["id"]

	for i, value := range Products {
		productID, err := strconv.Atoi(id)
		if err != nil || productID < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Invalid product ID."))
			return
		}

		if value.ProductID == productID {
			Products[i] = product
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 - Invalid product ID."))
}

// Updates the inventory value for the inventory item
func updateInventory(w http.ResponseWriter, r *http.Request) {

}
