package entity

// Product entity
type Product struct {
	ID       int `gorm:"primary_key, AUTO_INCREMENT"`
	Title    string
	Name     string
	Price    float64
	Quantity int
	Status   bool
}
