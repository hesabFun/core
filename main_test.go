package main

import (
	"github.com/stretchr/testify/assert"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestMainExecution(t *testing.T) {
	go main()
	time.Sleep(time.Second)
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatalf("Cannot make get: %v\n", err)
	}
	bodySb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading body: %v\n", err)
	}
	body := string(bodySb)
	fmt.Printf("Body: %v\n", body)
	var decodedResponse interface{}
	err = json.Unmarshal(bodySb, &decodedResponse)
	if err != nil {
		t.Fatalf("Cant decode response <%p> from server. Err: %v", bodySb, err)
	}
	assert.Equal(t, map[string]interface{}{"message": "Hello world!"}, decodedResponse,
		"Should return status:ok")
}
