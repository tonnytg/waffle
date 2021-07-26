package main

import (
	"database/sql"
	"github.com/tonnytg/waffle/adapters/db"
	"github.com/tonnytg/waffle/application"
	"log"
)

func main() {
	conn, err := sql.Open(db.DatabaseDriver, db.DataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	productDbAdapter := db.NewProductDb(conn)
	productService := application.NewProductService(productDbAdapter)

	product, err := productService.Create("foo bar", 25.0)
	if err != nil {
		log.Print("create: ", err)
	}
	// update quantity
	_, err = productService.SetQuantity(product, 11)
	if err != nil {
		log.Print("update: quantity", err)
	}
	// update status
	_, err = productService.Enable(product)
	if err != nil {
		log.Print("update: enabled", err)
	}
}
