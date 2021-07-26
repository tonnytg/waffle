package db

import (
	"database/sql"
	"github.com/tonnytg/waffle/application"
	"log"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, quantity, status from waffles where id=?")
	if err != nil {
		return nil, err
	}

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

	log.Println("create: db insert:", product.GetID())
	stmt, err := p.db.Prepare(`insert into waffles values ($1, $2, $3, $4, $5)`)
	// insert into waffles (id, name, price, quantity, status) values ('123e4567-e89b-12d3-a456-426614174000', 'test', 0.0, 1, 'disabled');
	if err != nil {
		log.Println("stmt: insert:", stmt)
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
	log.Println("create: success", product.GetID())
	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.db.Exec("update waffles set name = $1, price = $2, quantity = $3, status = $4 where id = $5",
		product.GetName(), product.GetPrice(), product.GetQuantity(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}
	log.Println("updated:", product.GetID())
	return product, nil
}
