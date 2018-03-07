package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorldRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	//JSON encode our body data
	jsonEncoded, _ := json.Marshal(gin.H{"message": "Hello world!"})

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(jsonEncoded), w.Body.String())
}
