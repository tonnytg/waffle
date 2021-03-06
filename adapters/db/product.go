package db

import (
	"database/sql"
	"github.com/tonnytg/waffle/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, quantity, status from waffles where id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var pID string
	rows, _ := p.db.Query("SELECT id FROM waffles where id = $1", product.GetID())
	for rows.Next() {
		rows.Scan(&pID)
	}
	if pID == "" {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`insert into waffles values ($1, $2, $3, $4, $5)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetQuantity(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.db.Exec("update waffles set name = $1, price = $2, quantity = $3, status = $4 where id = $5",
		product.GetName(), product.GetPrice(), product.GetQuantity(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}
	return product, nil
}
