package db_test

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tonnytg/waffle/adapters/db"
	"github.com/tonnytg/waffle/application"
	"log"
	"testing"
)

var Db *sql.DB
var err error

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE waffles (
		"id" string,
		"name" string,
		"price" float,
		"quantity" integer,
		"status" string
	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func createProduct(db *sql.DB) {
	insert := `insert into waffles values("123e4567-e89b-12d3-a456-426614174000","bar",0,0,"disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		t.Error(err)
	}

	resultUuid, err := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		t.Error("ID invalid format to parse to uuid")
	}
	if product.GetID() != resultUuid {
		t.Error("ID must be equal 123e4567-e89b-12d3-a456-426614174000")
	}
	if product.GetPrice() != 0.0 {
		t.Error("price must be equal 0")
	}
	if product.GetStatus() != "disabled" {
		t.Error("status must be equal disabled")
	}
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product := application.NewProduct()
	product.Name = "bar"
	product.Price = 0

	// test disabled product at create
	productResult, err := productDb.Save(product)
	if err != nil {
		t.Error(err)
	}
	if product.Name != productResult.GetName() {
		t.Error("product name must be equal productResult.GetName")
	}
	if product.Price != productResult.GetPrice() {
		t.Error("product price must be equal 0")
	}
	if product.Status != productResult.GetStatus() {
		t.Error("product status must be equal disabled")
	}

	// test enabled product at update
	product.Status = "enabled"
	product.Price = 25
	productResult, err = productDb.Save(product)
	if err != nil {
		t.Error(err)
	}
	if product.Name != productResult.GetName() {
		t.Error("product name must be equal productResult.GetName")
	}
	if product.Price != productResult.GetPrice() {
		t.Error("product price must be equal 25")
	}
	if product.Status != productResult.GetStatus() {
		t.Error("product status must be equal enabled")
	}

}
