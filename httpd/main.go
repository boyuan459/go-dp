package main

import (
	"dp/httpd/routes"
	"dp/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=dp password=pg_password sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := gin.Default()
	router.Use(middleware.DB(db))
	routes.Routes(router)
	router.Run()
}
