package main

import (
	"github.com/gin-gonic/gin/binding"
	"os"
	"testing"
	"upper.io/db.v3/mysql"
)

// This function is used to do setup before executing the test functions
func TestMain(m *testing.M) {
	binding.Validator = new(defaultValidator) // update validator to v9

	//Set Gin to Test Mode
	var DBError error
	MySql, DBError = mysql.Open(settings)
	if DBError != nil {
		//DBError
	}
	MySql.SetLogging(true)
	defer MySql.Close()

	// Run the other tests
	os.Exit(m.Run())
}
