package Host

import (
	"ShopProject/database"
	"ShopProject/errs"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func HandleUPCART(w http.ResponseWriter, r *http.Request) {

	session, err := cookie.Get(r, "ShopCookie")
	if err != nil {
		errs.Printer(err, "HandleUPCART1")
	}
	//id_str := r.URL.String()
	//id_str = strings.TrimSuffix(id_str, "/UPCART")
	//id_str = id_str[strings.LastIndexAny(id_str, "/"):]
	//id_str = strings.Replace(id_str, "/", "", 1)

	//id, err := strconv.Atoi(id_str)
	id := errs.CutID(r.URL.String(), "/UPCART")
	token := session.Values["token"]
	order := database.Order{Id: fmt.Sprint(token)}
	err = order.InsertOrder(database.Dv, database.DBCon, id)
	if err != nil {
		errs.Printer(err, "HandleUPCART2")
	}
}
func HandleLOWCART(w http.ResponseWriter, r *http.Request) {

	session, err := cookie.Get(r, "ShopCookie")
	if err != nil {
		errs.Printer(err, "HandleLOWCART1")
	}
	//id_str := r.URL.String()
	//id_str = strings.TrimSuffix(id_str, "/LOWCART")
	//id_str = id_str[strings.LastIndexAny(id_str, "/"):]
	//id_str = strings.Replace(id_str, "/", "", 1)

	//id, err := strconv.Atoi(id_str)
	id := errs.CutID(r.URL.String(), "/LOWCART")
	token := session.Values["token"]
	order := database.Order{Id: fmt.Sprint(token)}

	err = order.DeleteOneProduct(database.Dv, database.DBCon, id)
	if err != nil {
		errs.Printer(err, "HandleLOWCART2")
	}

}

func HandleCart(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/cart.html")
	if err != nil {
		errs.Printer(err, "HandleCart1")
	} else {
		session, err := cookie.Get(r, "ShopCookie")
		if err != nil {
			errs.Printer(err, "HandleCart2")
		}
		token := session.Values["token"]

		order := database.Order{Id: fmt.Sprint(token)}
		order.SelectOrder(database.Dv, database.DBCon)
		if err != nil {
			errs.Printer(err, "HandleCart3")
		} else {
			err := tmpl.Execute(w, order)
			if err != nil {
				errs.Printer(err, "HandleCart4")
			}
		}
	}
}

func HandlePay(w http.ResponseWriter, r *http.Request) {

	session, err := cookie.Get(r, "ShopCookie")
	if err != nil {
		errs.Printer(err, "HandlePay1")
	}
	token := session.Values["token"]
	log.Println("token in pay", token)
	order := database.Order{Id: fmt.Sprint(token)}
	err = order.Payment(database.Dv, database.DBCon)
	if err != nil {
		errs.Printer(err, "HandlePay2")
	}
}
func HandleDELETEFROMCART(w http.ResponseWriter, r *http.Request) {
	//id_str := r.URL.String()
	//id_str = strings.TrimSuffix(id_str, "/DELETEFROMCART")
	//id_str = id_str[strings.LastIndexAny(id_str, "/"):]
	//id_str = strings.Replace(id_str, "/", "", 1)
	//id, err := strconv.Atoi(id_str)
	id := errs.CutID(r.URL.String(), "/DELETEFROMCART")
	session, err := cookie.Get(r, "ShopCookie")
	if err != nil {
		errs.Printer(err, "HandleDELETEFROMCART1")
	}
	token := session.Values["token"]
	order := database.Order{Id: fmt.Sprint(token)}

	err = order.DeleteProduct(database.Dv, database.DBCon, id)
	if err != nil {
		errs.Printer(err, "HandleDELETEFROMCART2")
	}

}
func HandleCard(w http.ResponseWriter, r *http.Request) {
	session, err := cookie.Get(r, "ShopCookie")
	if err != nil {
		errs.Printer(err, "HandleCard1")
	}
	id := fmt.Sprintf("%v", session.Values["token"])
	log.Println("id is", id)
	var Orders = database.Order{Id: id}
	id_str := r.URL.String()
	id_str = strings.TrimSuffix(id_str, "/cart")
	id_str = id_str[strings.LastIndexAny(id_str, "/"):]
	id_str = strings.Replace(id_str, "/", "", 1)
	log.Println("id str is ", id_str)
	err, Shop := database.FindDataByID(database.Dv, database.DBCon, id_str)
	if err != nil {
		errs.Printer(err, "HandleCard2")
	}
	log.Println("shop is", Shop[0].Id)
	err = Orders.InsertOrder(database.Dv, database.DBCon, Shop[0].Id)
	if err != nil {
		errs.Printer(err, "HandleCard3")
	}
}
