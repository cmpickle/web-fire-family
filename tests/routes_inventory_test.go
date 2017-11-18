package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	// "../models"
	// "../routes"
	// "../app"

	"github.com/Xero67/web-fire-family/models"
	"github.com/Xero67/web-fire-family/routes"
	//"os"
)

func TestGetInventories(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for no so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/inventories", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"inventoryid", "quantity", "datelastupdated", "productid", "deleted"}).
		AddRow(1, 10, "11/17/2017", 1, 0).
		AddRow(2, 5, "11/16/2017", 2, 0).
		AddRow(3, 300, "11/15/2017", 3, 0)

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Inventory$").WillReturnRows(rows)
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"inventoryid":1,"quantity":10,"datelastupdated":"11/17/2017","productid":1},{"inventoryid":2,"quantity":5,"datelastupdated":"11/16/2017","productid":2},{"inventoryid":3,"quantity":300,"datelastupdated":"11/15/2017","productid":3}]`
	equal, err := AreEqualJSON(w.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestGetInventory(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/inventory/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"inventoryid", "quantity", "datelastupdated", "productid", "deleted"}).
		AddRow(1, 10, "11/17/2017", 1, 0)
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Inventory WHERE InventoryID = \\?$").WillReturnRows(rows)
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"inventoryid":1,"quantity":10,"datelastupdated":"11/17/2017","productid":1}]`
	equal, err := AreEqualJSON(w.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestGetInventoryInvalidID(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/inventory/8000", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Inventory WHERE InventoryID = \\?$").WillReturnError(fmt.Errorf("404 - Inventory not found"))

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestGetInventoryBySKU(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/inventorybysku/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"inventoryid", "quantity", "datelastupdated", "productid", "deleted"}).
		AddRow(1, 10, "11/17/2017", 1, 0)
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Inventory INNER JOIN Product ON Inventory.ProductID = Product.ProductID WHERE SKU = \\?$").WillReturnRows(rows)
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"inventoryid":1,"quantity":10,"datelastupdated":"11/17/2017","productid":1}]`
	equal, err := AreEqualJSON(w.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestGetInventoryNegativeID(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/product/-1", nil)
	if err != nil {
		t.Fatal(err)
	}

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	w := httptest.NewRecorder()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	expected := `400 - Invalid product ID.`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestCreateInventory(t *testing.T) {
	data := []byte(`{"productid":10,"productname":"Firefighter Stuff","notificationquantity":10,"color":"Tan","trimcolor":"Black","size":"size","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":10,"deleted":0}`)

	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/product/create", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO Product \\(ProductName, NotificationQuantity, Color, TrimColor, Size, Price, Dimensions, SKU\\) VALUES\\(\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?\\)").WithArgs("Firefighter Stuff", 10, "Tan", "Black", "size", 30.0, "3 1/2\" tall and 4 1/2\" long", 10).WillReturnResult(sqlmock.NewResult(10, 1))
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"ProductId": 10}`
	equal, err := AreEqualJSON(w.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestDeleteInventory(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("DELETE", "/product/delete/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"productid", "productname", "notificationquantity", "color", "trimcolor", "size", "price", "dimensions", "sku", "deleted"}).
		AddRow(2, "Swing", 10, "test", "test", "test", 1, "test", 1, 0)

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Product WHERE ProductID = \\?$").WillReturnRows(rows)
	mock.ExpectExec("^UPDATE Product SET Deleted = 1 WHERE ProductID = \\?").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"deleted": "true"}`
	equal, err := AreEqualJSON(w.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestDeleteInventoryNonExistant(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("DELETE", "/product/delete/8", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Product WHERE ProductID = \\?$").WillReturnError(fmt.Errorf("404 - Product not found"))

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	// Check the response body is what we expect.
	expected := `404 - Product not found`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestUpdateInventory(t *testing.T) {
	data := []byte(`{"productid":4,"productname":"Firefighter Stuff","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1}`)

	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("PUT", "/product/update/2", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"productid", "productname", "notificationquantity", "color", "trimcolor", "size", "price", "dimensions", "sku", "deleted"}).
		AddRow(2, "Swing", 10, "test", "test", "test", 1, "test", 1, 0)

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Product WHERE ProductID = \\?$").WillReturnRows(rows)
	mock.ExpectExec("^UPDATE Product SET ProductName = \\?, NotificationQuantity = \\?, Color = \\?, TrimColor = \\?, Size = \\?, Price = \\?, Dimensions = \\?, SKU = \\? WHERE ProductID = \\?$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusAccepted)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestUpdateInventoryInvalidID(t *testing.T) {
	data := []byte(`{"productid":4,"productname":"Firefighter Stuff","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1}`)

	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("PUT", "/product/update/8", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Product WHERE ProductID = \\?$").WillReturnError(fmt.Errorf("404 - Product not found")) //.WillReturnRows(rows)

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
