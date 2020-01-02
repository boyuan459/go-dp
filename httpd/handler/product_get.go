package handler

import (
	"dp/platform/product"

	"github.com/gin-gonic/gin"
)

// ProductGet test gin
// func ProductGet(products *product.List) gin.HandlerFunc {
func ProductGet(products product.Getter) gin.HandlerFunc {
	return func(context *gin.Context) {
		results := products.GetAll()
		context.JSON(200, results)
	}
}
