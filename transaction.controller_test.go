package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestInsertNewTransaction(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	jsonValue := `
{
	"product_id": 1,
	"title": "sale product 1",
	"amount": 10001,
	"type":"input",
	"date":"2018-02-18 12:12:12"
}
`

	req, _ := http.NewRequest("POST", "/v1/companies/1/transactions", bytes.NewBuffer([]byte(jsonValue)))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("JWT_TEST_TOKEN"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	//todo: test json schema
}
