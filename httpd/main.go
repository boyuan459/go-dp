package main

import (
	"dp/component"
	"dp/httpd/routes"
	"dp/middleware/auth"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	defer component.DB.Close()
	component.Router.Use(auth.JWT())
	component.Router.Use(auth.Casbin())
	routes.Routes()
	component.Router.Run()
}
