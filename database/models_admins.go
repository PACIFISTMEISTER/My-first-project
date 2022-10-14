package database

import (
	"ShopProject/errs"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

//модель админа
type Admin struct {
	Login    string `validate:"required,email"`
	Password string `validate:"required"`
}

//генерация  хэш пароля
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//сравнение пароля и хэш функции
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var validate *validator.Validate

//валидация введенных данных админа
func (a *Admin) Validation() error {

	validate = validator.New()
	err := validate.Struct(a)
	return err
}

//вставка нового админа из кода
func (admin *Admin) InsertData(dv string, conn string) error {
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Admin InsertData1")
	}
	admin.Password, err = HashPassword(admin.Password)
	if err != nil {
		errs.Printer(err, "Admin InsertData2")
	}
	insert := "insert into Admin(login,password) values ($1,$2);"
	_, err = db.Exec(insert, admin.Login, admin.Password)

	if err != nil {
		errs.Printer(err, "Admin InsertData3")
	}
	return nil

}

//проверка подлинности админа
func (admin *Admin) CheckAdmin(dv string, conn string) (error, bool) {
	ProperAdmin := Admin{"", ""}
	db, err := sql.Open(dv, conn)
	defer db.Close()
	if err != nil {
		errs.Printer(err, "Admin CheckAdmin1")
		return err, false
	}
	if err != nil {
		errs.Printer(err, "Admin CheckAdmin2")
	}
	sel := "select login,password from Admin"
	rows, err := db.Query(sel)

	if err != nil {
		errs.Printer(err, "Admin CheckAdmin3")
		return err, false
	}

	for rows.Next() {
		err = rows.Scan(&ProperAdmin.Login, &ProperAdmin.Password)

		if err != nil {
			errs.Printer(err, "Admin CheckAdmin4")
		} else {
			if CheckPasswordHash(admin.Password, ProperAdmin.Password) == true && (admin.Login == ProperAdmin.Login) {
				return nil, true
			} else {
				return nil, false
			}
		}
	}
	errs.Printer(err, "Admin CheckAdmin1")
	return err, false
}
