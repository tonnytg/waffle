package application

import "errors"

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLE = "disable"
	ENABLED = "enabled"
)

type Product struct {
	ID string
	Name string
	Price float64
	Status string
}

//func (p *Product) IsValid() (bool, error) {
//
//}
//
func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("product price must be greater than zero to enabled")
}
//
//func (p *Product) Disable() error {
//}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
