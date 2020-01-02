package router

import (
	"dp/httpd/router/product"

	"github.com/gin-gonic/gin"
)

// Router routes
type Router struct {
}

// New create a router
func New() *gin.Engine {
	router := gin.Default()
	product.Routes(router)

	return router
}
