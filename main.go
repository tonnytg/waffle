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

	_, err = productService.Create("foo bar", 25.0)
	if err != nil {
		log.Print("create: ", err)
	}
}
