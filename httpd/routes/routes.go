package routes

import (
	"dp/httpd/routes/misc"
	"dp/httpd/routes/product"
	"dp/httpd/routes/user"
)

// Routes list all routes
func Routes() {
	product.Routes()
	user.Routes()
	misc.Routes()
}
