package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/tonnytg/waffle/application"
	mock_application "github.com/tonnytg/waffle/application/mocks"
	"testing"
)

func TestProductService_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("foo")
	if err != nil {
		t.Error(err)
	}
	if result != product {
		t.Error("result must be equal to product")
	}
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("foo bar", 10)
	if err != nil {
		t.Error("service create broken")
	}

	if result != product {
		t.Error("result must be equal product")
	}
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)
	if err != nil {
		t.Error("product can't be enabled")
	}

	result, err = service.Disable(product)
	if err != nil {
		t.Error("product can't be enabled")
	}

	if result != product {
		t.Error("result must be equal product")
	}
}
