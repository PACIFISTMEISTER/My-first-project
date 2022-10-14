package database

import (
	"ShopProject/errs"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

//модель таблицы Shop
type Shop struct {
	Id          int
	Price       float64 `validate:"required,min=0"`
	Amount      int     `validate:"required,min=0"`
	Name        string  `validate:"required,gte=3,excludes=DELETE,excludes=PUTUP,excludes=PUTDOWN,excludes=AddNew"`
	Description string  `validate:"required,gte=5,excludes=DELETE,excludes=PUTUP,excludes=PUTDOWN,excludes=AddNew"`
	Category    string  `validate:"required,gte=5,excludes=DELETE,excludes=PUTUP,excludes=PUTDOWN,excludes=AddNew"`
	Picture     string  `validate:"required"`
}

func (shop *Shop) Validation() error {

	validate = validator.New()
	err := validate.Struct(shop)

	return err
}

//вставка данных
func (shop *Shop) InsertData(dv string, conn string) error {
	err := shop.Validation()
	if err != nil {
		errs.Printer(err, "Shop InsertData1")
		return err
	}
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Shop InsertData2")
	}
	insert := "insert into Shop(Price,Amount,Name,Description,Category,Picture)values ($1,$2,$3,$4,$5,$6);"
	_, err = db.Exec(insert, shop.Price, shop.Amount, shop.Name, shop.Description, shop.Category, shop.Picture)
	if err == nil {
		return nil
	}
	errs.Printer(err, "Shop InsertData3")
	return err

}

//удаление данных
func (shop *Shop) DeleteData(dv string, conn string) error {
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Shop DeleteData1")
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		errs.Printer(err, "Shop DeleteData1.5")
	}
	del := "delete from Shop where Id=$1;"
	_, err = tx.ExecContext(ctx, del, shop.Id)
	if err != nil {
		errs.Printer(err, "Shop DeleteData2")
		err := tx.Rollback()
		if err != nil {
			errs.Printer(err, "Shop DeleteData3")
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		errs.Printer(err, "Shop DeleteData3")
		return err
	}
	return nil

}

//вывод всех данных
func SelectDataForAdmin(dv string, conn string) (er error, temp []Shop) {

	NewShop := Shop{}
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Shop SelectDataForAdmin1")
		return err, nil
	}
	sel := "select Id,Price,Amount,Name,Description,Category,Picture from Shop ORDER BY Id;"
	rows, err := db.Query(sel)
	if err != nil {
		errs.Printer(err, "Shop SelectDataForAdmin2")
		return err, nil
	}
	for rows.Next() {
		err = rows.Scan(&NewShop.Id, &NewShop.Price, &NewShop.Amount, &NewShop.Name, &NewShop.Description, &NewShop.Category, &NewShop.Picture)
		if err != nil {
			errs.Printer(err, "Shop SelectDataForAdmin3")
		} else {
			temp = append(temp, NewShop)
		}
	}
	err = rows.Err()
	if err != nil {
		errs.Printer(err, "Shop SelectDataForAdmin4")
		return err, nil
	}
	defer rows.Close()
	return nil, temp

}

func SelectDataForUser(dv string, conn string) (er error, temp []Shop) {

	NewShop := Shop{}
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Shop SelectDataForUser1")
		return err, nil
	}
	sel := "select Id,Price,Amount,Name,Description,Category,Picture from Shop where Amount>0 ORDER BY Id;"
	rows, err := db.Query(sel)
	if err != nil {
		errs.Printer(err, "Shop SelectDataForUser2")
		return err, nil
	}
	for rows.Next() {
		err = rows.Scan(&NewShop.Id, &NewShop.Price, &NewShop.Amount, &NewShop.Name, &NewShop.Description, &NewShop.Category, &NewShop.Picture)
		if err != nil {
			errs.Printer(err, "Shop SelectDataForUser3")
		} else {
			temp = append(temp, NewShop)
		}
	}
	err = rows.Err()
	if err != nil {
		errs.Printer(err, "Shop SelectDataForUser4")
		return err, nil
	}
	defer rows.Close()
	return nil, temp

}

func FindDataByCategory(dv string, conn, Category string) (er error, temp []Shop) {

	NewShop := Shop{Category: Category}
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Shop FindDataByCategory1")
		return err, nil
	}
	sel := "select Id,Price,Amount,Name,Description,Picture from Shop where Category=$1 and Amount>0;"
	rows, err := db.Query(sel, NewShop.Category)
	if err != nil {
		errs.Printer(err, "Shop FindDataByCategory2")
		return err, nil
	}
	for rows.Next() {
		err = rows.Scan(&NewShop.Id, &NewShop.Price, &NewShop.Amount, &NewShop.Name, &NewShop.Description, &NewShop.Picture)
		if err != nil {
			errs.Printer(err, "Shop FindDataByCategory3")
		} else {
			temp = append(temp, NewShop)
		}
	}
	err = rows.Err()
	if err != nil {
		errs.Printer(err, "Shop FindDataByCategory4")
		return err, nil
	}
	defer rows.Close()
	return nil, temp

}

func FindDataByID(dv string, conn string, id string) (er error, temp []Shop) {

	NewShop := Shop{}
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Shop FindDataByID1")
		return err, nil
	}
	sel := "select Id,Price,Amount,Name,Description,Category,Picture from Shop where Id=$1;"
	rows, err := db.Query(sel, id)
	if err != nil {
		errs.Printer(err, "Shop FindDataByID2")
		return err, nil
	}
	for rows.Next() {
		err = rows.Scan(&NewShop.Id, &NewShop.Price, &NewShop.Amount, &NewShop.Name, &NewShop.Description, &NewShop.Category, &NewShop.Picture)
		if err != nil {
			errs.Printer(err, "Shop FindDataByID3")
		} else {
			temp = append(temp, NewShop)
		}
	}
	err = rows.Err()
	if err != nil {
		errs.Printer(err, "Shop FindDataByID4")
		return err, nil
	}
	defer rows.Close()
	return nil, temp

}

func (shop *Shop) UpdateData(dv string, conn string, i int) error {
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Shop UpdateData1")
		return err
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		errs.Printer(err, "Shop UpdateData1.5")
	}
	if i == 1 {

		updateUP := "update Shop set Amount=Amount+1 where id=$1;"
		_, err = tx.ExecContext(ctx, updateUP, shop.Id)
		if err != nil {
			errs.Printer(err, "Shop UpdateData1.7")
			err := tx.Rollback()
			if err != nil {
				errs.Printer(err, "Shop UpdateData2")
			}

			return err
		}
		err := tx.Commit()
		if err != nil {
			errs.Printer(err, "Shop UpdateData2.5")
		}
		return nil
	} else {
		updateUP := "update Shop set Amount=Amount-1 where id=$1;"
		_, err = tx.ExecContext(ctx, updateUP, shop.Id)
		if err != nil {
			errs.Printer(err, "Shop UpdateDat3")
			err := tx.Rollback()
			if err != nil {
				errs.Printer(err, "Shop UpdateData4")
			}

			return err
		}
		err := tx.Commit()
		if err != nil {
			errs.Printer(err, "Shop UpdateData2.5")
		}
		return nil
	}
}

func (shop *Shop) UpdateAllData(dv string, conn string) error {
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Shop UpdateAllData1")
	}
	err = shop.Validation()
	if err != nil {
		errs.Printer(err, "Shop UpdateAllData2")
		return err
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)

	updateUP := "update Shop set Price=$1,Amount=$2,Name=$3,Description=$4,Category=$5, Picture=$6 where id=$7;"
	_, err = tx.ExecContext(ctx, updateUP, shop.Price, shop.Amount, shop.Name, shop.Description, shop.Category, shop.Picture, shop.Id)
	if err == nil {
		err := tx.Commit()
		if err != nil {
			errs.Printer(err, "Shop UpdateAllData3")
		}

		return nil
	}

	errs.Printer(err, "Shop UpdateAllData4")
	err = tx.Rollback()
	if err != nil {
		errs.Printer(err, "Shop UpdateAllData5")
	}
	return err

}

type Category struct {
	Amount int
	Name   string
}

func SelectCategory(dv string, conn string) (er error, temp []Category) {

	Cat := Category{}
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Shop SelectCategory1")
		return err, nil
	}
	sel := "Select Category,Count(Category) from Shop group by Category"
	rows, err := db.Query(sel)
	if err != nil {
		errs.Printer(err, "Shop SelectCategory2")
		return err, nil
	}
	for rows.Next() {
		err = rows.Scan(&Cat.Name, &Cat.Amount)
		if err != nil {
			errs.Printer(err, "Shop SelectCategory3")
		} else {
			temp = append(temp, Cat)
		}
	}
	err = rows.Err()
	if err != nil {
		errs.Printer(err, "Shop SelectCategory4")
		return err, nil
	}
	defer rows.Close()
	return nil, temp

}

func (shop *Shop) ChangeAmount(dv string, conn string, i int) error {
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Shop ChangeAmount1")
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)

	updateUP := "update Shop set Amount=$1 where id=$2;"
	_, err = tx.ExecContext(ctx, updateUP, i, shop.Id)
	if err == nil {
		err := tx.Commit()
		if err != nil {
			errs.Printer(err, "Shop ChangeAmount2")
		}

		return nil
	}
	errs.Printer(err, "Shop ChangeAmount3")
	err = tx.Rollback()
	if err != nil {
		errs.Printer(err, "Shop ChangeAmount4")
	}
	return err

}
