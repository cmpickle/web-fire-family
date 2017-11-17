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

	// "../models"
	// "../routes"
	// "../app"
	"github.com/Xero67/web-fire-family/app"
	"github.com/Xero67/web-fire-family/models"
	"github.com/Xero67/web-fire-family/routes"
	//"os"
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

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
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

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
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

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Product WHERE ProductID = \\?$").WillReturnError(fmt.Errorf("404 - Product not found"))

	router := routes.InitRoutes(models.Env{db})

	router.ServeHTTP(w, req)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestGetProductNegativeID(t *testing.T) {
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

func TestCreateProduct(t *testing.T) {
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

func TestDeleteProduct(t *testing.T) {
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

func TestDeleteProductNonExistant(t *testing.T) {
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

func TestUpdateProduct(t *testing.T) {
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

func TestUpdateProductInvalidID(t *testing.T) {
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

func TestSettingYamlVar(t *testing.T) {
	//arrange
	var Dbdriver app.Dbdriver
	var host string
	var port int
	var dbuser string
	var dbpass string
	var database string
	host = "localhost"
	port = 3306
	dbuser = "test"
	dbpass = "letmein"
	database = "testDB"
	//act
	Dbdriver = Dbdriver.LoadSettings("../configtest.yml")
	//assert
	if host != Dbdriver.Host {
		t.Fatalf("Unable to find host")
	}

	if port != Dbdriver.Port {
		t.Fatalf("unable to find port")
	}

	if dbuser != Dbdriver.Dbuser {
		t.Fatalf("unable to find dbusername")
	}

	if dbpass != Dbdriver.Dbpass {
		t.Fatalf("Unable to find dbPass")
	}

	if database != Dbdriver.Database {
		t.Fatalf("Missing database string")
	}

}

func TestWebSettings(t *testing.T) {
	var port int = 8000
	var web app.Web

	web = web.LoadSettings("../configtest.yml")
	if port != web.Port {
		t.Fatalf("Missing Webport")
	}

}
