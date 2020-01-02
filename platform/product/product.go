package product

type Getter interface {
	GetAll() []Product
}

type Setter interface {
	Add(p Product)
}

type Product struct {
	Title string `json:"title"`
	Name  string `json:"name"`
}

type List struct {
	Products []Product
}

func New() *List {
	return &List{
		Products: []Product{},
	}
}

func (c *List) Add(p Product) {
	c.Products = append(c.Products, p)
}

func (c *List) GetAll() []Product {
	return c.Products
}
