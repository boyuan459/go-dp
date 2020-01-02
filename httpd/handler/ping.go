package handler

import "github.com/gin-gonic/gin"

// PingGet test gin
func PingGet() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"message": "hello world v2",
		})
	}
}
