package application_test

import (
	"github.com/google/uuid"
	"github.com/tonnytg/waffle/application"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "waffle"
	product.Status = application.DISABLED
	product.Price = 1
	product.Quantity = 1

	err := product.Enable()
	if err != nil {
		t.Error("product can't be enabled")
	}
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "waffle"
	product.Status = application.ENABLED
	product.Price = 10
	product.Quantity = 0

	err := product.Disable()
	if err != nil {
		t.Error("product can't be disabled")
	}
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "waffle"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.New()

	_, err := product.IsValid()
	if err != nil {
		t.Error("product is invalid")
	}
}
