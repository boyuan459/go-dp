package product

import (
	"dp/api/product"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	pRoute := router.Group("/product")
	{
		pRoute.GET("/", product.Get())
	}
}
