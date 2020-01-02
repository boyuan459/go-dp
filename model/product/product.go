package product

import (
	"dp/entities"

	"github.com/jinzhu/gorm"
)

// Product model
type Product struct {
	db *gorm.DB
}

// type List struct {
// 	Products []Product
// }

// func New() *List {
// 	return &List{
// 		Products: []Product{},
// 	}
// }

// func (c *List) Add(p Product) {
// 	c.Products = append(c.Products, p)
// }

// func (c *List) GetAll() []Product {
// 	return c.Products
// }

// New Product
func New(db *gorm.DB) *Product {
	return &Product{
		db: db,
	}
}

func (p *Product) Init() {
	p.db.AutoMigrate(&entities.Product{})
}

func (p *Product) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	p.db.Find(&products)
	return products, nil
}
