package main

import (
	"ShopProject/Host"
	"ShopProject/database"
	"ShopProject/errs"
	_ "github.com/lib/pq"
)

func main() {
	adm := database.Admin{"1mfree@mail.ru", "localhost"}
	adm.InsertData(database.Dv, database.DBCon)
	if err := Host.StartServer(); err != nil {
		errs.Printer(err, "error in server start")
	}

}
