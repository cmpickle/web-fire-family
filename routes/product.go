package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	//"github.com/Xero67/web-fire-family/models"
	"../models"
	"github.com/gorilla/mux"
)

// Returns all of the products stored in the database in JSON format
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	var rows *sql.Rows
	if rows, err = tx.Query("SELECT P.*, I.Quantity FROM Product P LEFT JOIN Inventory I ON P.ProductID = I.ProductID"); err != nil {
		return
	}

	prods := make([]*models.Product, 0)
	for rows.Next() {
		p := new(models.Product)
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.NotificationQuantity, &p.Color, &p.TrimColor, &p.Size, &p.Price, &p.Dimensions, &p.SKU, &p.Deleted, &p.Quantity)
		if err != nil {
			//More error handling
			fmt.Println("routes.go - getProducts - rows.Scan error")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	id := params["id"]
	found := -1

	productID, err := strconv.Atoi(id)
	if productID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product ID."))
		return
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	var rows *sql.Rows
	if rows, err = tx.Query("SELECT P.*, I.Quantity FROM Product P LEFT JOIN Inventory I ON P.ProductID = I.ProductID  WHERE P.ProductID = ?", id); err != nil {
		fmt.Println("routes.go - getProduct - tx.Query error selecting product id: " + id)
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer rows.Close()

	prods := make([]*models.Product, 0)
	for rows.Next() {
		p := new(models.Product)
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.NotificationQuantity, &p.Color, &p.TrimColor, &p.Size, &p.Price, &p.Dimensions, &p.SKU, &p.Deleted, &p.Quantity)
		if err != nil {
			//More error handling
			fmt.Println("2")
			fmt.Println(err)
		}
		if p.Deleted == 1 {
			prods = append(prods, p)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Product not found"))
			return
		}
		found = productID
		prods = append(prods, p)
	}
	if err = rows.Err(); err != nil {
		//Error handling
		fmt.Println("routes.go - getProduct - rows.Err()")
		fmt.Println(err)
	}

	//STILL NEED THIS FOR IF ITS NOT FOUND
	if found == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Product not found"))
		// json.NewEncoder(w).Encode("404 - Product not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prods)
}

// Returns a specific product from the database in JSON format
func getProductBySKU(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	sku := params["sku"]
	found := -1

	productSKU, err := strconv.Atoi(sku)
	if productSKU < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product SKU."))
		return
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	var rows *sql.Rows
	if rows, err = tx.Query("SELECT P.*, I.Quantity FROM Product P LEFT JOIN Inventory I ON P.ProductID = I.ProductID  WHERE P.SKU = ?", sku); err != nil {
		fmt.Println("routes.go - getProduct - tx.Query error selecting product sku: " + sku)
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer rows.Close()

	prods := make([]*models.Product, 0)
	for rows.Next() {
		p := new(models.Product)
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.NotificationQuantity, &p.Color, &p.TrimColor, &p.Size, &p.Price, &p.Dimensions, &p.SKU, &p.Deleted, &p.Quantity)
		if err != nil {
			//More error handling
			fmt.Println("2")
			fmt.Println(err)
		}
		if p.Deleted == 1 {
			prods = append(prods, p)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Product not found"))
			return
		}
		found = productSKU
		prods = append(prods, p)
	}
	if err = rows.Err(); err != nil {
		//Error handling
		fmt.Println("routes.go - getProduct - rows.Err()")
		fmt.Println(err)
	}

	//STILL NEED THIS FOR IF ITS NOT FOUND
	if found == -1 {
		w.WriteHeader(http.StatusNotFound)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(product)

	//new stuff
	//probs want to validate each required column
	if product.ProductName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product, please include a name, notification quantity, color, trim color, size, price, dimensions, and SKU"))
		return
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	res, err := tx.Exec("INSERT INTO Product (ProductName, NotificationQuantity, Color, TrimColor, Size, Price, Dimensions, SKU) VALUES(?,?,?,?,?,?,?,?)", product.ProductName, product.NotificationQuantity, product.Color, product.TrimColor, product.Size, product.Price, product.Dimensions, product.SKU)
	if err != nil {
		fmt.Println(err)
		fmt.Println("1")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product, please include a name, notification quantity, color, trim color, size, price, dimensions, and SKU"))
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
	w.Write([]byte("{\"ProductId\": " + lstId + "}"))
}

// Deletes the specified product from the database
// If the route logic were seperate from the DB logic, we could just call a getproductbyID method that is used
// by both
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	id := params["id"]
	found := -1

	productID, err := strconv.Atoi(id)
	if productID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product ID."))
		return
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	rows, err := tx.Query("SELECT * FROM Product WHERE ProductID = ?", id)
	if err != nil {
		//Error handlin
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Product not found"))
		fmt.Errorf("404 - Product not found")
		fmt.Println(err)
		return
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
		prods[0].Deleted = 1
		_, err := tx.Exec("UPDATE Product SET Deleted = 1 WHERE ProductID = ?", prods[0].ProductID)
		if err != nil {
			fmt.Println(err)
		}

	}
	w.Write([]byte(`{"deleted": "true"}`))
	w.WriteHeader(http.StatusOK)
}

func deleteProductBySKU(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	id := params["sku"]
	found := -1

	productID, err := strconv.Atoi(id)
	if productID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product ID."))
		return
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	rows, err := tx.Query("SELECT * FROM Product WHERE SKU = ?", id)
	if err != nil {
		//Error handlin
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Product not found"))
		fmt.Errorf("404 - Product not found")
		fmt.Println(err)
		return
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
		prods[0].Deleted = 1
		_, err := tx.Exec("UPDATE Product SET Deleted = 1 WHERE ProductID = ?", prods[0].ProductID)
		if err != nil {
			fmt.Println(err)
		}

	}
	w.Write([]byte(`{"deleted": "true"}`))
	w.WriteHeader(http.StatusOK)
}

// Updates the product
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	params := mux.Vars(r)
	id := params["id"]
	found := -1

	//new block
	productID, err := strconv.Atoi(id)
	if err != nil || productID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product ID."))
		return
	}
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	rows, err := tx.Query("SELECT * FROM Product WHERE ProductID = ?", id)
	if err != nil {
		//Error handlin
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Product not found"))
		fmt.Errorf("404 - Product not found")
		fmt.Println(err)
		return
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
		//need to do validation here
		res, err := tx.Exec("UPDATE Product SET ProductName = ?, NotificationQuantity = ?, Color = ?, TrimColor = ?, Size = ?, Price = ?, Dimensions = ?, SKU = ? WHERE ProductID = ?", product.ProductName, product.NotificationQuantity, product.Color, product.TrimColor, product.Size, product.Price, product.Dimensions, product.SKU, prods[0].ProductID)
		if err != nil {
			fmt.Println(err)
			fmt.Println("1")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Invalid product, please include a name, notification quantity, color, trim color, size, price, dimensions, and SKU"))
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
		fmt.Printf("update affected = %d\n", rowCnt)
		//Not sure what we want to return when success?
		w.WriteHeader(http.StatusAccepted)
	}
}

func updateProductBySKU(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	params := mux.Vars(r)
	id := params["sku"]
	found := -1

	//new block
	productID, err := strconv.Atoi(id)
	if err != nil || productID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid product ID."))
		return
	}
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	rows, err := tx.Query("SELECT * FROM Product WHERE SKU = ?", id)
	if err != nil {
		//Error handlin
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Product not found"))
		fmt.Errorf("404 - Product not found")
		fmt.Println(err)
		return
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
		//need to do validation here
		res, err := tx.Exec("UPDATE Product SET ProductName = ?, NotificationQuantity = ?, Color = ?, TrimColor = ?, Size = ?, Price = ?, Dimensions = ?, SKU = ? WHERE ProductID = ?", product.ProductName, product.NotificationQuantity, product.Color, product.TrimColor, product.Size, product.Price, product.Dimensions, product.SKU, prods[0].ProductID)
		if err != nil {
			fmt.Println(err)
			fmt.Println("1")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Invalid product, please include a name, notification quantity, color, trim color, size, price, dimensions, and SKU"))
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
		fmt.Printf("update affected = %d\n", rowCnt)
		//Not sure what we want to return when success?
		w.WriteHeader(http.StatusAccepted)
	}
}
