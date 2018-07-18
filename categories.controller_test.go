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

func TestInsertNewCategory(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	var category Categorise
	category.Name = "test category"
	category.Order = 1
	category.ParentId = 0
	jsonValue, _ := json.Marshal(category)

	req, _ := http.NewRequest("POST", "/v1/companies/1/categories", bytes.NewBuffer(jsonValue))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("JWT_TEST_TOKEN"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestInsertNewCategoryFailed(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	var category Categorise
	category.Name = ""
	category.Order = 0
	category.ParentId = 0
	jsonValue, _ := json.Marshal(category)

	req, _ := http.NewRequest("POST", "/v1/companies/1/categories", bytes.NewBuffer(jsonValue))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("JWT_TEST_TOKEN"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
