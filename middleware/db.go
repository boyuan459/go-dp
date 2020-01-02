package middleware

import (
	"dp/constants"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DB dependency injection to context
func DB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constants.DB, db)
		c.Next()
	}
}
