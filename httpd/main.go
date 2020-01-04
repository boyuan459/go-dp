package main

import (
	"dp/component"
	"dp/httpd/routes"
	"dp/middleware"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	defer component.DB.Close()
	component.Router.Use(middleware.JWT())
	routes.Routes()
	component.Router.Run()
}
