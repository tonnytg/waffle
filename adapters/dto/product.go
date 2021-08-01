package dto

import (
	"github.com/google/uuid"
	"github.com/tonnytg/waffle/application"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int32   `json:"quantity"`
	Status   string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = uuid.MustParse(p.ID)
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Quantity = p.Quantity
	product.Status = p.Status
	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}
	return product, nil
}
