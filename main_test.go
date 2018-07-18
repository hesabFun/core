package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMainExecution(t *testing.T) {
	c := make(chan int, 1)
	c <- 0
	os.Setenv("PORT", "7060")

	go main()
	time.Sleep(time.Second)
	resp, err := http.Get("http://localhost:7060/")
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

	close(c)
}
