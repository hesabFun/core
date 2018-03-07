package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

var settings = mysql.ConnectionURL{
	Database: os.Getenv("MYSQL_DATABASE"),
	Host:     os.Getenv("MYSQL_ADDRESS"),
	User:     os.Getenv("MYSQL_USERNAME"),
	Password: os.Getenv("MYSQL_PASSWORD"),
}

var MySql sqlbuilder.Database

func main() {

	var DBError error
	MySql, DBError = mysql.Open(settings)
	if DBError != nil {
		log.Fatal("MySQL Error: ", DBError)
	}
	MySql.SetLogging(false)
	defer MySql.Close()

	router := setupRouter()

	router.Run() // listen and serve on 0.0.0.0:8080
}
