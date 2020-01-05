package perm

import (
	"dp/model/casbin"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPerm(c *gin.Context) {
	var model casbin.Casbin
	err := c.BindJSON(&model)
	fmt.Println("model", model)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "invalid input",
		})
		return
	}
	ok := model.AddPolicy(model)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "add policy success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "add policy failure",
		})
	}
}
