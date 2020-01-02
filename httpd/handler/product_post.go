package handler

import (
	"dp/platform/product"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productPostRequest struct {
	Title string `json:"title"`
	Name  string `json:"name"`
}

// ProductPost test gin
// func ProductPost(products *product.List) gin.HandlerFunc {
func ProductPost(products product.Setter) gin.HandlerFunc {
	return func(context *gin.Context) {
		requestBody := productPostRequest{}
		context.Bind(&requestBody)
		p := product.Product{
			Title: requestBody.Title,
			Name:  requestBody.Name,
		}
		products.Add(p)
		fmt.Println("products", products)
		context.Status(http.StatusNoContent)
	}
}
