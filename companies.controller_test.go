package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestInsertNewCompany(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	jsonValue := `
{
	"name": "new company for test"
}
`

	req, _ := http.NewRequest("POST", "/v1/companies", bytes.NewBuffer([]byte(jsonValue)))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("JWT_TEST_TOKEN"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	//todo: test json schema
}

func TestGetAllCompanies(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v1/companies", nil)
	req.Header.Add("Authorization", "Bearer "+os.Getenv("JWT_TEST_TOKEN"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//todo: test json schema
}
