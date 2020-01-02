package product

import (
	"dp/entity"
	"dp/model"
)

type Getter interface {
	GetAll() []Product
}

type Setter interface {
	Add(p Product)
}

// Product model
type Product struct {
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

func FindAll() ([]entity.Product, error) {
	db, err := model.GetDB()
	if err != nil {
		return nil, err
	}
	var products []entity.Product
	db.Find(&products)
	return products, nil
}
