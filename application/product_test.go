package application_test

import (
	"github.com/tonnytg/waffle/application"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLE
	product.Price = 1

	err := product.Enable()
	if err != nil {
		t.Error()
	}
}