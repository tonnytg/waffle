package db

import "fmt"

const DatabaseDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Password = "password"
const DbName = "waffledb"
const TableName = "waffles"

var Port = "5432"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
