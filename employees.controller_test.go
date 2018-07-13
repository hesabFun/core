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

func TestAddNewEmployee(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	var employee Employees
	employee.UserId = 2
	employee.Type = "technical"
	jsonValue, _ := json.Marshal(employee)

	req, _ := http.NewRequest("POST", "/v1/companies/1/employees", bytes.NewBuffer(jsonValue))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("JWT_TEST_TOKEN"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestGetAllEmployees(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v1/companies/1/employees", nil)
	req.Header.Add("Authorization", "Bearer "+os.Getenv("JWT_TEST_TOKEN"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
