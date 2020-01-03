package product

import (
	"dp/component"
	"dp/httpd/handler/product"
)

func Routes() {
	pRoute := component.Router.Group("/product")
	{
		pRoute.GET("/", product.Get())
	}
}
