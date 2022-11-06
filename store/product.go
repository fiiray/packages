package store

var standardTaxRate = NewTaxRate(0.25, 20)

type Product struct {
	Name, Category string
	price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{
		Name:     name,
		Category: category,
		price:    price,
	}
}

func (p *Product) Price() float64 {
	return standardTaxRate.calcTax(p)
}

func (p *Product) SetPrice(newPrice float64) {
	p.price = newPrice
}
