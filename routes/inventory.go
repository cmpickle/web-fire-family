package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	//"github.com/Xero67/web-fire-family/models"
	"../models"
	"github.com/gorilla/mux"
)

func getInventories(w http.ResponseWriter, r *http.Request) {
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
	if rows, err = tx.Query("SELECT * FROM Inventory"); err != nil {
		return
	}

	inv := make([]*models.Inventory, 0)
	for rows.Next() {
		i := new(models.Inventory)
		err := rows.Scan(&i.InventoryID, &i.Quantity, &i.DateLastUpdated, &i.ProductID, &i.Deleted)
		if err != nil {
			//More error handling
			fmt.Println("routes.go - getInventory - rows.Scan error")
			fmt.Println(err)
		}
		if i.Deleted == 0 {
			inv = append(inv, i)
		}
	}
	if err = rows.Err(); err != nil {
		//Error handling
		fmt.Println("3")
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(inv)
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	found := -1

	//new stuff can easily change to work off of SKU
	inventoryID, err := strconv.Atoi(id)
	if inventoryID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid inventory ID."))
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
	if rows, err = tx.Query("SELECT * FROM Inventory WHERE InventoryID = ?", id); err != nil {
		fmt.Println("inventory.go - getInventory - tx.Query error selecting inventory id: " + id)
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer rows.Close()

	inv := make([]*models.Inventory, 0)
	for rows.Next() {
		i := new(models.Inventory)
		err := rows.Scan(&i.InventoryID, &i.Quantity, &i.DateLastUpdated, &i.ProductID, &i.Deleted)
		if err != nil {
			//More error handling
			fmt.Println("2")
			fmt.Println(err)
		}
		if i.Deleted == 1 {
			inv = append(inv, i)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Product not found"))
			return
		}
		found = inventoryID
		inv = append(inv, i)
	}
	if err = rows.Err(); err != nil {
		//Error handling
		fmt.Println("inventory.go - getInventory - rows.Err()")
		fmt.Println(err)
	}

	//STILL NEED THIS FOR IF ITS NOT FOUND
	if found == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Product not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inv)
}

// Updates the inventory value for the inventory item
func updateInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	found := -1

	//new stuff can easily change to work off of SKU
	inventoryID, err := strconv.Atoi(id)
	if inventoryID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid inventory ID."))
		return
	}

	var inventory models.Inventory
	err = json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		fmt.Println(err)
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
	if rows, err = tx.Query("SELECT * FROM Inventory WHERE InventoryID = ?", id); err != nil {
		fmt.Println("inventory.go - getInventory - tx.Query error selecting inventory id: " + id)
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer rows.Close()

	inv := make([]*models.Inventory, 0)
	for rows.Next() {
		i := new(models.Inventory)
		err := rows.Scan(&i.InventoryID, &i.Quantity, &i.DateLastUpdated, &i.ProductID, &i.Deleted)
		if err != nil {
			//More error handling
			fmt.Println("2")
			fmt.Println(err)
		}
		if i.Deleted == 1 {
			inv = append(inv, i)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Product not found"))
			return
		}
		found = inventoryID
		inv = append(inv, i)
	}
	if err = rows.Err(); err != nil {
		//Error handling
		fmt.Println("inventory.go - getInventory - rows.Err()")
		fmt.Println(err)
	}

	//STILL NEED THIS FOR IF ITS NOT FOUND
	if found == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Product not found"))
		return
	}

	if rows, err = tx.Query("UPDATE Inventory SET InventoryID = ?, Quantity = ?, DateLastUpdated = ?, ProductID = ?, Deleted = ? WHERE InventoryID = ?", inv[0].InventoryID, params["quantity"], time.Now(), inv[0].ProductID, inv[0].Deleted, id); err != nil {
		fmt.Println("inventory.go - getInventory - tx.Query error selecting inventory id: " + id)
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer rows.Close()

	w.WriteHeader(http.StatusOK)
}

func incrementInventory(w http.ResponseWriter, r *http.Request) {

}

func decrementInventory(w http.ResponseWriter, r *http.Request) {

}
