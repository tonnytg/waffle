package application

import "errors"

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error ) {

	if id == "" {
		return nil, errors.New("id is required")
	}

	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}