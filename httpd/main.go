package main

import (
	"dp/httpd/router"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=pg_password sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := router.New()
	router.Run()
}
