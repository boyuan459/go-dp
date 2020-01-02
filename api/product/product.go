package product

import (
	"github.com/gin-gonic/gin"
)

// Get get product
func Get() gin.HandlerFunc {
	return func(context *gin.Context) {
		// results, _ := product.FindAll()
		results := []string{"test", "go"}
		context.JSON(200, results)
	}
}
