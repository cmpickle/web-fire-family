package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Xero67/web-fire-family/routes"
)

func TestGetProducts(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for no so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router := routes.InitRoutes()

	routes.Products = nil

	routes.Products = append(routes.Products, routes.Product{ProductID: 1, ProductName: "Firefighter Wallet", InventoryScanningID: 1, Color: "Tan", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	routes.Products = append(routes.Products, routes.Product{ProductID: 2, ProductName: "Firefighter Apron", InventoryScanningID: 2, Color: "Tan", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	routes.Products = append(routes.Products, routes.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", InventoryScanningID: 3, Color: "Tan", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"productid":1,"productname":"Firefighter Wallet","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1},{"productid":2,"productname":"Firefighter Apron","inventoryscanningid":2,"color":"Tan","size":"One Size Fits All","price":29,"dimensions":"31\" tall and 26\" wide and ties around a waist up to 54\"","sku":2},{"productid":3,"productname":"Firefighter Baby Outfit","inventoryscanningid":3,"color":"Tan","size":"Newborn","price":39.99,"dimensions":"Waist-14\", Length-10\"","sku":3}]`
	equal, err := AreEqualJSON(w.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestGetProduct(t *testing.T) {
	// Create a request to pass to our handler. We don't ahve any query parameters for no so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/product/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router := routes.InitRoutes()

	routes.Products = nil

	routes.Products = append(routes.Products, routes.Product{ProductID: 1, ProductName: "Firefighter Wallet", InventoryScanningID: 1, Color: "Tan", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	routes.Products = append(routes.Products, routes.Product{ProductID: 2, ProductName: "Firefighter Apron", InventoryScanningID: 2, Color: "Tan", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	routes.Products = append(routes.Products, routes.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", InventoryScanningID: 3, Color: "Tan", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"productid":1,"productname":"Firefighter Wallet","inventoryscanningid":1,"color":"Tan","price":30,"dimensions":"3 1/2\" tall and 4 1/2\" long","sku":1}`
	equal, err := AreEqualJSON(w.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestGetProductInvalidID(t *testing.T) {
	// Create a request to pass to our handler. We don't ahve any query parameters for no so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/product/8", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router := routes.InitRoutes()

	routes.Products = nil

	routes.Products = append(routes.Products, routes.Product{ProductID: 1, ProductName: "Firefighter Wallet", InventoryScanningID: 1, Color: "Tan", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	routes.Products = append(routes.Products, routes.Product{ProductID: 2, ProductName: "Firefighter Apron", InventoryScanningID: 2, Color: "Tan", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	routes.Products = append(routes.Products, routes.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", InventoryScanningID: 3, Color: "Tan", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `""`
	equal, err := AreEqualJSON(w.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestGetProductNegativeID(t *testing.T) {
	// Create a request to pass to our handler. We don't ahve any query parameters for no so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/product/-1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router := routes.InitRoutes()

	routes.Products = nil

	routes.Products = append(routes.Products, routes.Product{ProductID: 1, ProductName: "Firefighter Wallet", InventoryScanningID: 1, Color: "Tan", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	routes.Products = append(routes.Products, routes.Product{ProductID: 2, ProductName: "Firefighter Apron", InventoryScanningID: 2, Color: "Tan", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	routes.Products = append(routes.Products, routes.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", InventoryScanningID: 3, Color: "Tan", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `""`
	equal, err := AreEqualJSON(w.Body.String(), expected)
	if !equal {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
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

	router := routes.InitRoutes()

	routes.Products = nil

	routes.Products = append(routes.Products, routes.Product{ProductID: 1, ProductName: "Firefighter Wallet", InventoryScanningID: 1, Color: "Tan", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	routes.Products = append(routes.Products, routes.Product{ProductID: 2, ProductName: "Firefighter Apron", InventoryScanningID: 2, Color: "Tan", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	routes.Products = append(routes.Products, routes.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", InventoryScanningID: 3, Color: "Tan", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

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

	router := routes.InitRoutes()

	routes.Products = nil

	routes.Products = append(routes.Products, routes.Product{ProductID: 1, ProductName: "Firefighter Wallet", InventoryScanningID: 1, Color: "Tan", Price: 30, Dimensions: "3 1/2\" tall and 4 1/2\" long", SKU: 1})
	routes.Products = append(routes.Products, routes.Product{ProductID: 2, ProductName: "Firefighter Apron", InventoryScanningID: 2, Color: "Tan", Size: "One Size Fits All", Price: 29, Dimensions: "31\" tall and 26\" wide and ties around a waist up to 54\"", SKU: 2})
	routes.Products = append(routes.Products, routes.Product{ProductID: 3, ProductName: "Firefighter Baby Outfit", InventoryScanningID: 3, Color: "Tan", Size: "Newborn", Price: 39.99, Dimensions: "Waist-14\", Length-10\"", SKU: 3})

	router.ServeHTTP(w, req)

	router.ServeHTTP(w2, req2)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
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
