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

	result, err := service.Get("abc")
	if err != nil {
		t.Error(err)
	}
	if result != product {
		t.Error("result must be equal to product")
	}
}