package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterNewUser(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	var employee struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	employee.Mobile = "09111111113"
	employee.Password = "12345678"
	employee.Name = "test register new user"

	jsonValue, _ := json.Marshal(employee)
	req, _ := http.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}
