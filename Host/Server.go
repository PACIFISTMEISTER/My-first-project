package Host

import (
	"ShopProject/database"
	"ShopProject/errs"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"strings"
	"time"
)

var cookie *sessions.CookieStore

func init() {
	cookie = sessions.NewCookieStore([]byte("ShopCookie"))

}

// StartServer старт сервера, а так же обслуживание опредленных адресов

//TODO изменил с двойных скобок на одинарные
func StartServer() error {
	r := mux.NewRouter()
	r.Use(TokenMiddleware)
	r.Use(URLMiddleware)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	r.HandleFunc("/hello", HandleHello)                                                      // showdata1
	r.HandleFunc("/Category", HandleAllCategory)                                             // showdata1
	r.HandleFunc("/registration", CheckRegister)                                             // other1
	r.HandleFunc("/hello/{category}", HandleCategory)                                        //showdata1
	r.HandleFunc("/hello/{category}/{id}", HandleId)                                         //showdata1
	r.HandleFunc("/hello/{category}/{id}/DELETE", HandleDelete)                              // admin1
	r.HandleFunc("/hello/{category}/{id}/PUTUP", HandlePutUP)                                // admin1
	r.HandleFunc("/hello/{category}/{id}/PUTDOWN", HandlePutDOWN)                            // admin1
	r.HandleFunc("/hello/{product.Category}/{product.Id}/UPDATE", HandleUpdate)              //admin1
	r.HandleFunc("/hello/{product.Category}/{product.Id }/CHECKPRODUCT", HandleCHECKPRODUCT) //showdata1
	r.HandleFunc("/hello/{product.Category}/{product.Id}/static/{picture}", HandleImage)     //other1
	r.HandleFunc("/Admin", HandleAdmin)                                                      //admin1
	r.HandleFunc("/hello/{product.Category}/{$product.Id}/cart", HandleCard)                 // cart
	r.HandleFunc("/cart", HandleCart)                                                        // cart
	r.HandleFunc("/pay", HandlePay)                                                          // cart
	r.HandleFunc("/hello/{category}/{id}/LOWCART", HandleLOWCART)                            //cart
	r.HandleFunc("/hello/{category}/{id}/UPCART", HandleUPCART)                              // cart
	r.HandleFunc("/hello/{category}/{id}/DELETEFROMCART", HandleDELETEFROMCART)              //cart
	r.HandleFunc("/AddNew", HandleAddNew)                                                    //admin1
	s := http.Server{
		Addr:         Address,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
	}
	return s.ListenAndServe()
}

func CheckRegister(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			errs.Printer(err, "CheckRegister1")
		} else {
			PossibleAdmin := database.Admin{}
			PossibleAdmin.Login = r.FormValue("login")
			PossibleAdmin.Password = r.FormValue("password")

			err, IsAdmin := PossibleAdmin.CheckAdmin(database.Dv, database.DBCon)
			if err != nil {
				errs.Printer(err, "CheckRegister1")
			}

			if IsAdmin == true {
				token, err := database.GenerateTokenForAdmin()
				err, _ = database.ParseToken(token)
				if err != nil {
					errs.Printer(err, "CheckRegister2")
				}

				if err != nil {
					errs.Printer(err, "CheckRegister3")
				}
				session, err := cookie.Get(r, "ShopCookie")
				if err != nil {
					errs.Printer(err, "CheckRegister4")
				}
				session.Values["token"] = token
				session.Save(r, w)

				http.Redirect(w, r, "/Admin", 301)
			} else {
				if err != nil {
					errs.Printer(err, "CheckRegister5")
				}
				http.Redirect(w, r, "/hello", 301)
			}

		}
	} else {
		session, err := cookie.Get(r, "ShopCookie")
		if err != nil {
			errs.Printer(err, "CheckRegister6")
		}
		token := session.Values["token"]

		err, s := database.ParseToken(fmt.Sprint(token))

		if s == "True" {
			token, err = database.GenerateTokenForUser()
			if err != nil {
				errs.Printer(err, "CheckRegister7")
			}
			session.Values["token"] = token
			err := session.Save(r, w)
			if err != nil {
				errs.Printer(err, "CheckRegister8")
			}

		}
		tmpl, err := template.ParseFiles("static/registration.html")

		if err != nil {
			errs.Printer(err, "CheckRegister9")
		} else {
			err := tmpl.Execute(w, nil)
			if err != nil {
				errs.Printer(err, "CheckRegister10")
			}
		}
	}
}

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := cookie.Get(r, "ShopCookie")
		session.Options = &sessions.Options{
			MaxAge: 86400,
		}
		if err != nil {
			errs.Printer(err, "TokenMiddleware1")
		}
		if session.Values["token"] == nil {
			token, err := database.GenerateTokenForUser()
			if err != nil {
				errs.Printer(err, "TokenMiddleware2")
			}

			session.Values["token"] = token
			err = session.Save(r, w)
			if err != nil {
				errs.Printer(err, "TokenMiddleware3")
			}
		}

		err, _ = database.ParseToken(fmt.Sprint(session.Values["token"]))
		if err != nil {
			errs.Printer(err, "TokenMiddleware4")
		}
		next.ServeHTTP(w, r)
	})
}

func URLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AdminRequire := [...]string{"DELETE", "PUTUP", "PUTDOWN", "UPDATE", "AddNew", "Admin"}
		url := r.URL.String()
		url = url[strings.LastIndexAny(url, "/")+1:]
		IsUrlForbitten := false
		for _, word := range AdminRequire {
			if word == url {
				IsUrlForbitten = true
				goto exit
			}
		}
	exit:
		if IsUrlForbitten == true {
			session, err := cookie.Get(r, "ShopCookie")
			
			if err != nil {
				errs.Printer(err, "URLMiddleware1")
			}
			token := session.Values["token"]

			err, s := database.ParseToken(fmt.Sprint(token))
			if err != nil {
				errs.Printer(err, "URLMiddleware2")
			}
			if s == "True" {
				next.ServeHTTP(w, r)
			} else {
				http.Redirect(w, r, "/hello", 403)
			}
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
