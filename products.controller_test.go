package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestInsertNewProduct(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	var product Products
	product.Name = "test category"
	product.CategoryId = 1
	product.Description = "product description"
	product.Price = 10000
	jsonValue, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/v1/companies/1/products", bytes.NewBuffer(jsonValue))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("JWT_TEST_TOKEN"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestInsertNewProductFailed(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	var product Products
	product.Name = ""
	product.CategoryId = 0
	product.Description = ""
	product.Price = 0
	jsonValue, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/v1/companies/1/products", bytes.NewBuffer(jsonValue))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("JWT_TEST_TOKEN"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestGetAllProducts(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v1/companies/1/products", nil)
	req.Header.Add("Authorization", "Bearer "+os.Getenv("JWT_TEST_TOKEN"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
