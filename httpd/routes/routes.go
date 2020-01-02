package routes

import (
	"dp/httpd/routes/product"

	"github.com/gin-gonic/gin"
)

// Routes list all routes
func Routes(router *gin.Engine) {
	product.Routes(router)
}
