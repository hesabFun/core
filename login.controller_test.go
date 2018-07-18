package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	jsonValue, _ := json.Marshal(map[string]string{
		"mobile":   "09111111111",
		"password": "12345678",
	})

	req, _ := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestLoginFailed(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	jsonValue, _ := json.Marshal(map[string]string{
		"mobile":   "",
		"password": "",
	})

	req, _ := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestLoginWrongPassword(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	jsonValue, _ := json.Marshal(map[string]string{
		"mobile":   "09111111111",
		"password": "1111111111",
	})

	req, _ := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
