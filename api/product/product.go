package product

import (
	"dp/constants"
	"dp/model/product"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

// Get get product
func Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		// product.Init()
		db := c.MustGet(constants.DB).(*gorm.DB)
		productModel := product.New(db)
		results, _ := productModel.FindAll()
		// results := []string{"test", "go"}
		c.JSON(200, results)
	}
}
