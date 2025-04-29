package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var CurrentUserAc int

func ServeLoginPage(w http.ResponseWriter, r *http.Request) {
	// temp, err := template.ParseFiles("HTML/Login.html")
	// if err != nil {
	// 	http.Error(w, "unable to load the login page", http.StatusInternalServerError)
	// 	return
	// }
	temp := LoadhtmlPage(w, "Login.html")
	temp.Execute(w, nil)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/bank/login", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	accStr := r.FormValue("userAccNoFromLogin")
	username := r.FormValue("usernameFromLogin")
	pass := r.FormValue("passwordFromLogin")
	query := "select cus_name,password from CustomersList where acc_no = ?"

	accInt, err := strconv.Atoi(accStr)
	if err != nil {
		http.Redirect(w, r, "/bank/main", http.StatusSeeOther)
		return
	}

	var (
		name     string
		password string
	)

	type PageData struct {
		ShowMessage bool
	}

	err = Db.QueryRow(query, accInt).Scan(&name, &password)
	if err == sql.ErrNoRows {

		http.Redirect(w, r, "/bank/main", http.StatusSeeOther)
	} else if username == name && pass == password {
		http.SetCookie(w, &http.Cookie{
			Name:     "account_number",
			Value:    accStr,
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
		})

		http.Redirect(w, r, "/bank/serveIndex", http.StatusSeeOther)
	} else {
		templ, _ := template.ParseFiles("HTML/Login.html")
		data := PageData{ShowMessage: true}
		templ.Execute(w, data)
	}
}
