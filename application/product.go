package application

import (
	"errors"
	"github.com/google/uuid"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() uuid.UUID
	GetName() string
	GetStatus() string
	GetPrice() float64
	GetQuantity() int32
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWrite interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWrite
}

const (
	DISABLED = "disable"
	ENABLED  = "enabled"
)

type Product struct {
	ID       uuid.UUID
	Name     string
	Price    float64
	Status   string
	Quantity int32
}

func NewProduct() *Product {
	product := Product{
		ID:     uuid.New(),
		Status: DISABLED,
	}
	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("status must be enabled or disabled")
	}

	if p.Price <= 0 {
		return false, errors.New("price must be greater or equal zero")
	}

	if p.Quantity < 0 {
		return false, errors.New("quantity must be greater zero")
	}

	if len(p.Name) <= 0 {
		return false, errors.New("name cannot be empty")
	}

	if len(p.ID) <= 0 {
		return false, errors.New("id cannot be empty")
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 && p.Quantity > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("product price and quantity must be greater than zero to enabled")
}

func (p *Product) Disable() error {
	if p.Quantity == 0 {
		p.Status = DISABLED
		p.Price = 0
		return nil
	}
	return errors.New("product quantity equal 0 to disabled")
}

func (p *Product) GetID() uuid.UUID {
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

func (p *Product) GetQuantity() int32 {
	return p.Quantity
}
