package entities

import "github.com/jinzhu/gorm"

// Product entity
type Product struct {
	gorm.Model
	Title    string
	Name     string
	Price    float64
	Quantity int
	Status   bool
}
