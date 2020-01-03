package product

import (
	"dp/component"
	"dp/entities"
)

func init() {
	component.DB.AutoMigrate(&entities.Product{})
}

// FindAll find all products
func FindAll() ([]entities.Product, error) {
	var products []entities.Product
	component.DB.Find(&products)
	return products, nil
}
