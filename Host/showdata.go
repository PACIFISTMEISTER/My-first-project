package Host

import (
	"ShopProject/database"
	"ShopProject/errs"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("static/hello.html")

	if err != nil {
		errs.Printer(err, "HandleHello1")
	} else {
		err, allData := database.SelectDataForUser(database.Dv, database.DBCon)
		if err != nil {
			errs.Printer(err, "HandleHello2")
		} else {

			if err != nil {
				errs.Printer(err, "HandleHello3")
			}

			err := tmpl.Execute(w, allData)
			if err != nil {
				errs.Printer(err, "HandleHello4")
			}
		}
	}
}

func HandleAllCategory(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("static/Category.html")
	if err != nil {
		errs.Printer(err, "HandleAllCategory1")
	} else {
		err, allData := database.SelectCategory(database.Dv, database.DBCon)
		if err != nil {
			errs.Printer(err, "HandleAllCategory2")
		} else {
			err := tmpl.Execute(w, allData)
			if err != nil {
				errs.Printer(err, "HandleAllCategory3")
			}
		}
	}
}

func HandleId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.String()
	id = id[strings.LastIndexAny(id, "/"):]
	id = strings.Replace(id, "/", "", 1)
	tmpl, err := template.ParseFiles("static/hello.html")
	if err != nil {
		errs.Printer(err, "HandleId1")
	} else {
		err, allData := database.FindDataByID(database.Dv, database.DBCon, id)
		if err != nil {
			errs.Printer(err, "HandleId2")
		} else {
			err := tmpl.Execute(w, allData)
			if err != nil {
				errs.Printer(err, "HandleId3")
			}
		}
	}
}

func HandleCategory(w http.ResponseWriter, r *http.Request) {
	cat := r.URL.String()
	cat = strings.Replace(cat, "/hello/", "", 1)
	cat, err := url.QueryUnescape(cat)
	if err != nil {
		errs.Printer(err, "HandleCategory1")
	}

	tmpl, err := template.ParseFiles("static/hello.html")
	if err != nil {
		errs.Printer(err, "HandleCategory2")
	} else {
		err, allData := database.FindDataByCategory(database.Dv, database.DBCon, cat)
		if err != nil {
			errs.Printer(err, "HandleCategory3")
		} else {
			err := tmpl.Execute(w, allData)
			if err != nil {
				errs.Printer(err, "HandleCategory4")
			}
		}
	}
}

func HandleCHECKPRODUCT(w http.ResponseWriter, r *http.Request) {

	id := errs.CutID(r.URL.String(), "/CHECKPRODUCT")
	err, NewData := database.FindDataByID(database.Dv, database.DBCon, fmt.Sprint(id))
	if err != nil {
		errs.Printer(err, "HandleCHECKPRODUCT1")
	}
	tmpl, err := template.ParseFiles("static/ProductInfo.html")
	if err != nil {
		errs.Printer(err, "HandleCHECKPRODUCT2")
	}

	err = tmpl.Execute(w, NewData)
	if err != nil {
		errs.Printer(err, "HandleCHECKPRODUCT3")
	}
}
