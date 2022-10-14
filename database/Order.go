package database

import (
	"ShopProject/errs"
	"database/sql"
	"fmt"
	"log"
)

type ProductForOrder struct {
	Product         Shop
	AmountOfProduct int
}

type Order struct {
	Id       string
	Products []ProductForOrder
}

func (order *Order) InsertOrder(dv string, conn string, id int) error {
	{
		db, err := sql.Open(dv, conn)
		defer db.Close()
		if err != nil {
			errs.Printer(err, "Order InsertOrder1")
		}
		insert := "insert into order_details values($1,$2);"
		_, err = db.Exec(insert, order.Id, id)
		if err == nil {
			return nil
		}
		errs.Printer(err, "Order InsertOrder2")
		return err

	}
}

func (order *Order) DeleteOrder(dv string, conn string) error {
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Order DeleteOrder1")
	}
	insert := "delete from order_details where id_of_person=$1"
	_, err = db.Exec(insert, order.Id)
	if err == nil {
		return nil
	}
	errs.Printer(err, "Order DeleteOrder2")
	return err

}

func (order *Order) DeleteOneProduct(dv string, conn string, id int) error {
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Order DeleteOneProduct1")
		return err
	}
	del := "WITH u AS ( SELECT distinct on (id_of_product, id_of_person) id_of_product, id_of_person, ctid     FROM order_details) DELETE FROM order_details WHERE id_of_product=$1 AND id_of_person=$2 AND ctid IN (SELECT ctid FROM u)"
	_, err = db.Exec(del, id, order.Id)
	if err == nil {
		return nil
	}
	errs.Printer(err, "Order DeleteOneProduct2")
	return err

}

func (order *Order) DeleteProduct(dv string, conn string, id int) error {
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Order DeleteProduct1")
		return err
	}
	sel := " DELETE FROM order_details WHERE id_of_product=$1 AND id_of_person=$2 "
	_, err = db.Exec(sel, id, order.Id)
	if err == nil {
		return nil
	}
	errs.Printer(err, "Order DeleteProduct2")
	return err

}

func (order *Order) Payment(dv string, conn string) error {
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Order Payment1")
		return err
	}
	product_id := 0
	amount := 0
	var dif []int
	Pr := ProductForOrder{}
	sel := "select id_of_product, count(id_of_product) from order_details where id_of_person= $1 group by (id_of_product);"
	rows, err := db.Query(sel, order.Id)

	for rows.Next() {
		if err != nil {
			errs.Printer(err, "Order Payment2")
		} else {
			err = rows.Scan(&product_id, &amount)
			err, data := FindDataByID(Dv, DBCon, fmt.Sprint(product_id))
			if err != nil {
				errs.Printer(err, "Order Payment3")
			}

			Pr.AmountOfProduct = amount
			Pr.Product = data[0]
			order.Products = append(order.Products, Pr)
			dif = append(dif, data[0].Amount-amount)
		}
		err = rows.Err()
		if err != nil {

			errs.Printer(err, "Order Payment4")
		}
		defer rows.Close()

	}

	for _, d := range dif {
		if d < 0 {
			order.DeleteOrder(Dv, DBCon)
			return nil
		}
	}

	for i := 0; i < len(order.Products); i++ {

		err := order.Products[i].Product.ChangeAmount(Dv, DBCon, dif[i])
		if err != nil {
			errs.Printer(err, "Order Payment4")
		}
	}

	del := "delete from order_details where id_of_person=$1"
	_, err = db.Exec(del, order.Id)
	if err == nil {
		return nil
	}
	errs.Printer(err, "Order Payment5")
	return err

}
func (order *Order) SelectOrder(dv string, conn string) {
	product_id := 0
	amount := 0
	Pr := ProductForOrder{}
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		log.Println(err)
	}

	sel := "select id_of_product, count(id_of_product) from order_details where id_of_person= $1 group by (id_of_product);"
	rows, err := db.Query(sel, order.Id)
	if err != nil {
		errs.Printer(err, "Order SelectOrder1")
	}
	for rows.Next() {
		err = rows.Scan(&product_id, &amount)
		if err != nil {
			errs.Printer(err, "Order SelectOrder2")
		} else {

			err, data := FindDataByID(Dv, DBCon, fmt.Sprint(product_id))
			if err != nil {
				errs.Printer(err, "Order SelectOrder3")
			}

			Pr.AmountOfProduct = amount
			Pr.Product = data[0]
			order.Products = append(order.Products, Pr)

		}
	}
	err = rows.Err()
	if err != nil {
		errs.Printer(err, "Order SelectOrder4")
	}
	defer rows.Close()
}
