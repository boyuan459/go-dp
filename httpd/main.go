package main

import (
	"dp/httpd/handler"
	"dp/platform/product"

	"github.com/gin-gonic/gin"
)

func main() {
	products := product.New()
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/product", handler.ProductGet(products))
	r.POST("/product", handler.ProductPost(products))
	r.Run()
}
