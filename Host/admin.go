package Host

import (
	"ShopProject/database"
	"ShopProject/errs"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/succsesfullog.html")
	if err != nil {
		errs.Printer(err, "HandleAdmin1")
	}
	err, allData := database.SelectDataForAdmin(database.Dv, database.DBCon)
	if err != nil {
		errs.Printer(err, "HandleAdmin2")
	}
	err = tmpl.Execute(w, allData)
	if err != nil {
		errs.Printer(err, "HandleAdmin3")
	}
}

func HandlePutUP(w http.ResponseWriter, r *http.Request) {
	//idStr := r.URL.String()
	//idStr = strings.TrimSuffix(idStr, "/PUTUP")
	//idStr = idStr[strings.LastIndexAny(idStr, "/"):]
	//idStr = strings.Replace(idStr, "/", "", 1)

	//id, err := strconv.Atoi(idStr)
	id := errs.CutID(r.URL.String(), "/PUTUP")

	tempData := database.Shop{Id: id}

	err := tempData.UpdateData(database.Dv, database.DBCon, 1)
	if err != nil {
		errs.Printer(err, "HandlePutUP2")
	}
}

func HandlePutDOWN(w http.ResponseWriter, r *http.Request) {
	//idStr := r.URL.String()
	//idStr = strings.TrimSuffix(idStr, "/PUTDOWN")
	//idStr = idStr[strings.LastIndexAny(idStr, "/"):]
	//idStr = strings.Replace(idStr, "/", "", 1)

	//id, err := strconv.Atoi(idStr)
	id := errs.CutID(r.URL.String(), "/PUTDOWN")

	tempData := database.Shop{Id: id}

	err := tempData.UpdateData(database.Dv, database.DBCon, 0)
	if err != nil {
		errs.Printer(err, "HandlePutDOWN2")
	}
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	//idStr := r.URL.String()
	//idStr = strings.TrimSuffix(idStr, "/DELETE")
	//idStr = idStr[strings.LastIndexAny(idStr, "/"):]
	//idStr = strings.Replace(idStr, "/", "", 1)
	//id, err := strconv.Atoi(idStr)
	id := errs.CutID(r.URL.String(), "/DELETE")

	tempData := database.Shop{Id: id}

	err := tempData.DeleteData(database.Dv, database.DBCon)
	if err != nil {
		errs.Printer(err, "HandleDelete2")
	}
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseMultipartForm(40000000000)
		if err != nil {
			errs.Printer(err, "HandleUpdate1")
		} else {
			//id_str := r.URL.String()
			//id_str = strings.TrimSuffix(id_str, "/UPDATE")
			//id_str = id_str[strings.LastIndexAny(id_str, "/"):]
			//id_str = strings.Replace(id_str, "/", "", 1)
			//id, _ := strconv.Atoi(id_str)
			id := errs.CutID(r.URL.String(), "/UPDATE")

			NewShop := database.Shop{Id: id}
			NewShop.Name = r.FormValue("NewName")
			if NewShop.Name == "" {

				http.Redirect(w, r, "", 501)
			}

			NewShop.Description = r.FormValue("NewDescription")
			if NewShop.Description == "" {

				http.Redirect(w, r, "", 501)
			}
			NewShop.Category = r.FormValue("NewCategory")
			if NewShop.Category == "" {

				http.Redirect(w, r, "", 501)
			}
			NewShop.Price, err = strconv.ParseFloat(r.FormValue("NewPrice"), 64)
			if err != nil {
				errs.Printer(err, "HandleUpdate1.5")
			}
			if NewShop.Price == 0 {

				http.Redirect(w, r, "", 501)
			}
			NewShop.Amount, err = strconv.Atoi(r.FormValue("NewAmount"))
			if err != nil {
				errs.Printer(err, "HandleUpdate1.8")
			}
			if NewShop.Amount == 0 {

				http.Redirect(w, r, "", 501)
			}
			file, handle, err := r.FormFile("NewFile")
			if err != nil {
				errs.Printer(err, "HandleUpdate2")
				err, shop := database.FindDataByID(database.Dv, database.DBCon, fmt.Sprint(id))
				if err != nil {
					errs.Printer(err, "HandleUpdate2.5")
				}
				NewShop.Picture = shop[0].Picture
			} else {
				ext := strings.ToLower(path.Ext(handle.Filename))
				if ext != ".jpeg" {
					log.Println("bad format")
				}

				err = os.Mkdir("/", 0777)
				if err != nil {
					errs.Printer(err, "HandleUpdate3")
				}
				saveFile, err := os.OpenFile("./images/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					errs.Printer(err, "HandleUpdate4")
				}
				_, err = io.Copy(saveFile, file)
				if err != nil {
					errs.Printer(err, "HandleUpdate5")
				}
				defer file.Close()
				defer saveFile.Close()
				NewShop.Picture = handle.Filename
			}
			err = NewShop.UpdateAllData(database.Dv, database.DBCon)
			if err != nil {
				errs.Printer(err, "HandleUpdate6")
			}

			http.Redirect(w, r, r.URL.String(), 301)

		}
	} else {
		//id_str := r.URL.String()
		//id_str = strings.TrimSuffix(id_str, "/UPDATE")
		//id_str = id_str[strings.LastIndexAny(id_str, "/"):]
		//id_str = strings.Replace(id_str, "/", "", 1)
		id := errs.CutID(r.URL.String(), "/UPDATE")
		tmpl, err := template.ParseFiles("static/UpdateData.html")
		if err != nil {
			errs.Printer(err, "HandleUpdate7")
		} else {
			err, allData := database.FindDataByID(database.Dv, database.DBCon, fmt.Sprint(id))
			if err != nil {
				errs.Printer(err, "HandleUpdate8")
			} else {
				err := tmpl.Execute(w, allData)
				if err != nil {
					errs.Printer(err, "HandleUpdate8")
				}
			}
		}
	}
}

func HandleAddNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseMultipartForm(400000000)
		if err != nil {
			errs.Printer(err, "HandleAddNew1")
		} else {

			NewShop := database.Shop{}
			NewShop.Name = r.FormValue("NewName")
			log.Println(NewShop.Name)
			if NewShop.Name == "" {

				http.Redirect(w, r, "", 501)
			}
			NewShop.Description = r.FormValue("NewDescription")
			if NewShop.Description == "" {

				http.Redirect(w, r, "", 501)
			}
			NewShop.Category = r.FormValue("NewCategory")
			if NewShop.Category == "" {

				http.Redirect(w, r, "", 501)
			}
			NewShop.Price, err = strconv.ParseFloat(r.FormValue("NewPrice"), 64)

			if err != nil {
				errs.Printer(err, "HandleAddNew2")
			}
			if NewShop.Price == 0 {

				http.Redirect(w, r, "", 501)
			}
			NewShop.Amount, err = strconv.Atoi(r.FormValue("NewAmount"))
			if err != nil {
				errs.Printer(err, "HandleAddNew3")
			}
			if NewShop.Amount == 0 {

				http.Redirect(w, r, "", 501)
			}
			file, handle, err := r.FormFile("NewFile")
			if err != nil {
				errs.Printer(err, "HandleAddNew4")
				http.Redirect(w, r, "", 501)
			} else {
				ext := strings.ToLower(path.Ext(handle.Filename))
				if ext != ".jpeg" && ext != ".png" {
					log.Println("bad format")
					//http.Redirect(w, r, "", 501)
				}

				err = os.Mkdir("/", 0777)
				if err != nil {
					errs.Printer(err, "HandleAddNew5")
				}
				saveFile, err := os.OpenFile("./images/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					errs.Printer(err, "HandleAddNew6")
				}
				_, err = io.Copy(saveFile, file)
				if err != nil {
					errs.Printer(err, "HandleAddNew7")
				}
				defer file.Close()
				defer saveFile.Close()
				NewShop.Picture = handle.Filename
			}

			err = NewShop.InsertData(database.Dv, database.DBCon)
			if err != nil {
				errs.Printer(err, "HandleAddNew8")
			}
			if err != nil {
				errs.Printer(err, "HandleAddNew9")
			}
			http.Redirect(w, r, r.URL.String(), 301)

		}
	} else {
		tmpl, err := template.ParseFiles("static/AddNew.html")
		if err != nil {
			errs.Printer(err, "HandleAddNew10")
		} else {
			err := tmpl.Execute(w, nil)
			if err != nil {
				errs.Printer(err, "HandleAddNew11")
			}
		}
	}

}
