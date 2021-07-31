package main

import (
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"github.com/tonnytg/waffle/adapters/db"
	"github.com/tonnytg/waffle/adapters/web/server"
	"github.com/tonnytg/waffle/application"
	"log"
)

func main() {

	f := flag.String("run", "http", "create\nhttp")
	// Check some Args different of default
	flag.Parse()

	if *f == "http" {
		webserver := server.MakeNewWebserver()
		webserver.Service = &application.ProductService{}
		log.Println("webserver: WebServer has been started")
		webserver.Serve()
	}

	if *f == "create" {
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
		product, err = productService.Get("133e4567-e89b-12d3-a456-426614174000")
		if err != nil {
			log.Println("update: enabled", err)
		}
		log.Println("waffle:", product)
	}
}
