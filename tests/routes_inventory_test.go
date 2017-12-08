package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/Xero67/web-fire-family/models"
	"github.com/Xero67/web-fire-family/routes"
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
	rows := sqlmock.NewRows([]string{"inventoryid", "quantity", "datelastupdated", "productid", "deleted", "sku"}).
		AddRow(1, 10, "11/17/2017", 0, 1, 1).
		AddRow(2, 5, "11/16/2017", 0, 2, 2).
		AddRow(3, 300, "11/15/2017", 0, 3, 3)

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT I.(.+), P.SKU FROM Inventory I INNER JOIN Product P ON P.ProductID = I.ProductID$").WillReturnRows(rows)
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"inventoryid":1,"quantity":10,"datelastupdated":"11/17/2017","productid":1,"sku":1},{"inventoryid":2,"quantity":5,"datelastupdated":"11/16/2017","productid":2,"sku":2},{"inventoryid":3,"quantity":300,"datelastupdated":"11/15/2017","productid":3,"sku":3}]`
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
	req, err := http.NewRequest("GET", "/inventory/4", nil)
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
	rows := sqlmock.NewRows([]string{"inventoryid", "quantity", "datelastupdated", "deleted", "productid", "SKU"}).
		AddRow(4, 10, "11/17/2017", 0, 1, 4)
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT I.(.+), P.SKU FROM Inventory I INNER JOIN Product P ON I.ProductID = P.ProductID WHERE P.SKU = \\?$").WillReturnRows(rows)
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"inventoryid":4,"quantity":10,"datelastupdated":"11/17/2017","productid":1,"sku":4}]`
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
	mock.ExpectQuery("^SELECT I.(.+), P.SKU FROM Inventory I INNER JOIN Product P ON I.ProductID = P.ProductID WHERE P.SKU = \\?$").WillReturnError(fmt.Errorf("404 - Inventory not found"))

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

// func TestGetInventoryBySKU(t *testing.T) {
// 	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
// 	req, err := http.NewRequest("GET", "/inventorybysku/1", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	w := httptest.NewRecorder()

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	// before we actually execute our api function, we need to expect required DB actions
// 	rows := sqlmock.NewRows([]string{"inventoryid", "quantity", "datelastupdated", "productid", "deleted"}).
// 		AddRow(1, 10, "11/17/2017", 1, 0)
// 	mock.ExpectBegin()
// 	mock.ExpectQuery("^SELECT (.+) FROM Inventory INNER JOIN Product ON Inventory.ProductID = Product.ProductID WHERE SKU = \\?$").WillReturnRows(rows)
// 	mock.ExpectCommit()

// 	router := routes.InitRoutes(models.Env{db})

// 	router.ServeHTTP(w, req)

// 	// Check the status code is what we expect.
// 	if status := w.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	// Check the response body is what we expect.
// 	expected := `[{"inventoryid":1,"quantity":10,"datelastupdated":"11/17/2017","productid":1}]`
// 	equal, err := AreEqualJSON(w.Body.String(), expected)
// 	if !equal {
// 		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
// 	}

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expections: %s", err)
// 	}
// }

func TestGetInventoryNegativeID(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/inventory/-1", nil)
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
	expected := `400 - Invalid product SKU.`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestUpdateInventory(t *testing.T) {
	data := []byte(`{"inventoryid":1,"quantity":10,"datelastupdated":"11/17/2017","productid":1,"deleted":1}`)

	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("PUT", "/inventory/update/1/50", bytes.NewBuffer(data))
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
	rows := sqlmock.NewRows([]string{"inventoryid", "quantity", "datelastupdated", "productid", "deleted"}).
		AddRow(1, 50, "11/17/2017", 1, 0)

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT I.(.+), P.SKU FROM Inventory I INNER JOIN Product P ON P.ProductID = I.ProductID WHERE P.SKU = \\?$").WillReturnRows(rows)
	mock.ExpectExec("^UPDATE Inventory SET InventoryID = \\?, Quantity = \\?, DateLastUpdated = \\?, Deleted = \\?, ProductID = \\? WHERE InventoryID = \\?$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestUpdateInventoryInvalidID(t *testing.T) {
	data := []byte(`{"inventoryid":1,"quantity":10,"datelastupdated":"11/17/2017","productid":1,"deleted":1}`)

	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("PUT", "/inventory/update/800/1", bytes.NewBuffer(data))
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
	mock.ExpectQuery("^SELECT I.(.+), P.SKU FROM Inventory I INNER JOIN Product P ON P.ProductID = I.ProductID WHERE P.SKU = \\?$").WillReturnError(fmt.Errorf("404 - Inventory not found"))

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

func TestIncrementInventory(t *testing.T) {
	data := []byte(`{"inventoryid":1,"quantity":11,"datelastupdated":"11/17/2017","productid":1,"deleted":1}`)

	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("PUT", "/inventory/increment/1", bytes.NewBuffer(data))
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
	rows := sqlmock.NewRows([]string{"inventoryid", "quantity", "datelastupdated", "productid", "deleted"}).
		AddRow(1, 11, "11/17/2017", 1, 0)

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT I.(.+), P.SKU FROM Inventory I INNER JOIN Product P on P.ProductID = I.ProductID WHERE P.SKU = \\?$").WillReturnRows(rows)
	mock.ExpectExec("^UPDATE Inventory SET InventoryID = \\?, Quantity = \\?, DateLastUpdated = \\?, Deleted = \\?, ProductID = \\? WHERE InventoryID = \\?$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestDecrementInventory(t *testing.T) {
	data := []byte(`{"inventoryid":1,"quantity":9,"datelastupdated":"11/17/2017","productid":1,"deleted":1}`)

	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("PUT", "/inventory/decrement/1", bytes.NewBuffer(data))
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
	rows := sqlmock.NewRows([]string{"inventoryid", "quantity", "datelastupdated", "productid", "deleted"}).
		AddRow(1, 9, "11/17/2017", 1, 0)

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT I.(.+), P.SKU FROM Inventory I INNER JOIN Product P on P.ProductID = I.ProductID WHERE P.SKU = \\?$").WillReturnRows(rows)
	mock.ExpectExec("^UPDATE Inventory SET InventoryID = \\?, Quantity = \\?, DateLastUpdated = \\?, Deleted = \\?, ProductID = \\? WHERE InventoryID = \\?$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
