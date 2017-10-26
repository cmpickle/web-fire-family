package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/Xero67/web-fire-family/models"
	"github.com/Xero67/web-fire-family/routes"
)

func TestGetProducts(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for no so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/product", nil)
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
		AddRow(1, "Firefighter Wallet", 10, "Tan", "Black", "size", 30, "3 1/2\" tall and 4 1/2\" long", 1, 0).
		AddRow(2, "Firefighter Apron", 20, "Tan", "Black", "One Size Fits All", 29, "31\" tall and 26\" wide and ties around a waist up to 54\"", 2, 0).
		AddRow(3, "Firefighter Baby Outfit", 13, "Tan", "Black", "Newborn", 39.99, "Waist-14\", Length-10\"", 3, 0)

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Product$").WillReturnRows(rows)
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"productid":1,"productname":"Firefighter Wallet","notificationquantity":10,"color":"Tan","trimcolor":"Black","size":"size","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1},{"productid":2,"productname":"Firefighter Apron","notificationquantity":20,"color":"Tan","trimcolor":"Black","size":"One Size Fits All","price":29,"dimensions":"31\" tall and 26\" wide and ties around a waist up to 54\"","sku":2},{"productid":3,"productname":"Firefighter Baby Outfit","notificationquantity":13,"color":"Tan","trimcolor":"Black","size":"Newborn","price":39.99,"dimensions":"Waist-14\", Length-10\"","sku":3}]`
	equal, err := AreEqualJSON(w.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestGetProduct(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/product/1", nil)
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
		AddRow(1, "Firefighter Wallet", 10, "Tan", "Black", "size", 30, "3 1/2\" tall and 4 1/2\" long", 1, 0)
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Product WHERE ProductID = \\?$").WillReturnRows(rows)
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"productid":1,"productname":"Firefighter Wallet","notificationquantity":10,"color":"Tan","trimcolor":"Black","size":"size","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1}]`
	equal, err := AreEqualJSON(w.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestGetProductInvalidID(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/product/8000", nil)
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
	// rows := sqlmock.NewRows([]string{"productid", "productname", "notificationquantity", "color", "trimcolor", "size", "price", "dimensions", "sku", "deleted"}).
	// 	AddRow(1, "Firefighter Wallet", 10, "Tan", "Black", "size", 30, "3 1/2\" tall and 4 1/2\" long", 1, 0)
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Product WHERE ProductID = \\?$").WillReturnError(fmt.Errorf("404 - Product not found"))
	mock.ExpectCommit()

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

func TestGetProductNegativeID(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/product/-1", nil)
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
		AddRow(1, "Firefighter Wallet", 10, "Tan", "Black", "size", 30, "3 1/2\" tall and 4 1/2\" long", 1, 0)
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Product$").WillReturnRows(rows)
	mock.ExpectCommit()

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

func TestCreateProduct(t *testing.T) {
	data := []byte(`{"productid":4,"productname":"Firefighter Stuff","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1}`)

	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/product/create", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	req2, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	w2 := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT * FROM Product").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	routes.Products = nil

	routes.Products = append(routes.Products, models.Product{ProductID: 1, ProductName: "Firefighter Wallet", NotificationQuantity: 10, Color: "Tan", TrimColor: "Black", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	routes.Products = append(routes.Products, models.Product{ProductID: 2, ProductName: "Firefighter Apron", NotificationQuantity: 20, Color: "Tan", TrimColor: "Black", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	routes.Products = append(routes.Products, models.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", NotificationQuantity: 13, Color: "Tan", TrimColor: "Black", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.ServeHTTP(w, req)

	router.ServeHTTP(w2, req2)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"productid":1,"productname":"Firefighter Wallet","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1},{"productid":2,"productname":"Firefighter Apron","inventoryscanningid":2,"color":"Tan","size":"One Size Fits All","price":29,"dimensions":"31\" tall and 26\" wide and ties around a waist up to 54\"","sku":2},{"productid":3,"productname":"Firefighter Baby Outfit","inventoryscanningid":3,"color":"Tan","size":"Newborn","price":39.99,"dimensions":"Waist-14\", Length-10\"","sku":3},{"productid":4,"productname":"Firefighter Stuff","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1}]`
	equal, err := AreEqualJSON(w2.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w2.Body.String(), expected)
	}
}

func TestDeleteProduct(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("DELETE", "/product/delete/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	w2 := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT * FROM Product").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	routes.Products = nil

	routes.Products = append(routes.Products, models.Product{ProductID: 1, ProductName: "Firefighter Wallet", NotificationQuantity: 10, Color: "Tan", TrimColor: "Black", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	routes.Products = append(routes.Products, models.Product{ProductID: 2, ProductName: "Firefighter Apron", NotificationQuantity: 20, Color: "Tan", TrimColor: "Black", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	routes.Products = append(routes.Products, models.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", NotificationQuantity: 13, Color: "Tan", TrimColor: "Black", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.ServeHTTP(w, req)

	router.ServeHTTP(w2, req2)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"productid":1,"productname":"Firefighter Wallet","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1},{"productid":3,"productname":"Firefighter Baby Outfit","inventoryscanningid":3,"color":"Tan","size":"Newborn","price":39.99,"dimensions":"Waist-14\", Length-10\"","sku":3}]`
	equal, err := AreEqualJSON(w2.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w2.Body.String(), expected)
	}
}

func TestDeleteProductNonExistant(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("DELETE", "/product/delete/8", nil)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	w2 := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT * FROM Product").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	routes.Products = nil

	routes.Products = append(routes.Products, models.Product{ProductID: 1, ProductName: "Firefighter Wallet", NotificationQuantity: 10, Color: "Tan", TrimColor: "Black", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	routes.Products = append(routes.Products, models.Product{ProductID: 2, ProductName: "Firefighter Apron", NotificationQuantity: 20, Color: "Tan", TrimColor: "Black", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	routes.Products = append(routes.Products, models.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", NotificationQuantity: 13, Color: "Tan", TrimColor: "Black", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.ServeHTTP(w, req)

	router.ServeHTTP(w2, req2)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	expected := `[{"productid":1,"productname":"Firefighter Wallet","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1},{"productid":2,"productname":"Firefighter Apron","inventoryscanningid":2,"color":"Tan","size":"One Size Fits All","price":29,"dimensions":"31\" tall and 26\" wide and ties around a waist up to 54\"","sku":2},{"productid":3,"productname":"Firefighter Baby Outfit","inventoryscanningid":3,"color":"Tan","size":"Newborn","price":39.99,"dimensions":"Waist-14\", Length-10\"","sku":3}]`
	equal, err := AreEqualJSON(w2.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w2.Body.String(), expected)
	}
}

func TestUpdateProduct(t *testing.T) {
	data := []byte(`{"productid":4,"productname":"Firefighter Stuff","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1}`)

	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("PUT", "/product/update/2", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	req2, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	w2 := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT * FROM Product").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	routes.Products = nil

	routes.Products = append(routes.Products, models.Product{ProductID: 1, ProductName: "Firefighter Wallet", NotificationQuantity: 10, Color: "Tan", TrimColor: "Black", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	routes.Products = append(routes.Products, models.Product{ProductID: 2, ProductName: "Firefighter Apron", NotificationQuantity: 20, Color: "Tan", TrimColor: "Black", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	routes.Products = append(routes.Products, models.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", NotificationQuantity: 13, Color: "Tan", TrimColor: "Black", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.ServeHTTP(w, req)

	router.ServeHTTP(w2, req2)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"productid":1,"productname":"Firefighter Wallet","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1},{"productid":4,"productname":"Firefighter Stuff","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1},{"productid":3,"productname":"Firefighter Baby Outfit","inventoryscanningid":3,"color":"Tan","size":"Newborn","price":39.99,"dimensions":"Waist-14\", Length-10\"","sku":3}]`
	equal, err := AreEqualJSON(w2.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w2.Body.String(), expected)
	}
}

func TestUpdateProductInvalidID(t *testing.T) {
	data := []byte(`{"productid":4,"productname":"Firefighter Stuff","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1}`)

	// Create a request to pass to our handler. We don't have any query parameters for now so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("PUT", "/product/update/8", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	req2, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	w2 := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT * FROM Product").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := routes.InitRoutes(models.Env{db})

	routes.Products = nil

	routes.Products = append(routes.Products, models.Product{ProductID: 1, ProductName: "Firefighter Wallet", NotificationQuantity: 10, Color: "Tan", TrimColor: "Black", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	routes.Products = append(routes.Products, models.Product{ProductID: 2, ProductName: "Firefighter Apron", NotificationQuantity: 20, Color: "Tan", TrimColor: "Black", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	routes.Products = append(routes.Products, models.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", NotificationQuantity: 13, Color: "Tan", TrimColor: "Black", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.ServeHTTP(w, req)

	router.ServeHTTP(w2, req2)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	expected := `[{"productid":1,"productname":"Firefighter Wallet","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1},{"productid":2,"productname":"Firefighter Apron","inventoryscanningid":2,"color":"Tan","size":"One Size Fits All","price":29,"dimensions":"31\" tall and 26\" wide and ties around a waist up to 54\"","sku":2},{"productid":3,"productname":"Firefighter Baby Outfit","inventoryscanningid":3,"color":"Tan","size":"Newborn","price":39.99,"dimensions":"Waist-14\", Length-10\"","sku":3}]`
	equal, err := AreEqualJSON(w2.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w2.Body.String(), expected)
	}
}

func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}
