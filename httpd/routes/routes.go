package routes

import (
	"dp/component"
	nonauth "dp/httpd/routes/misc"
	"dp/httpd/routes/perm"
	"dp/httpd/routes/product"
	"dp/httpd/routes/user"
	"dp/middleware/auth"
)

// Routes list all routes
func Routes() {
	nonauth.Routes()
	component.Router.Use(auth.JWT())
	component.Router.Use(auth.Casbin())
	product.Routes()
	user.Routes()
	perm.Routes()
}
