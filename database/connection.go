package database

import "fmt"

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "royalflush228"
	Dbname   = "postgres"
)

//connection
var DBCon = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
	Host, Port, User, Password, Dbname)

//driver
var Dv = "postgres"
