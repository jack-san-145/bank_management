package main

import (
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type balanceStruct struct {
	PasswordCorrect bool
	PasswordWrong   bool
	Balance         int
}

func ServeBalancePage(w http.ResponseWriter, r *http.Request) {
	if !checkForCookie(w, r) {
		return
	}
	templ := LoadhtmlPage(w, "BalanceInquiry.html")
	templ.Execute(w, nil)
}

func BalanceHandler(w http.ResponseWriter, r *http.Request) {

	if !checkForCookie(w, r) {
		return
	}
	var (
		account_password string
		account_balance  int
	)
	templ, _ := template.ParseFiles("HTML/BalanceInquiry.html")

	r.ParseForm()
	pass := r.FormValue("user_password")
	data := balanceStruct{false, false, 0}
	_, _, account_password, _, _, account_balance, _ = FindUserFromCookie(w, r)
	if pass == account_password {
		data.Balance = account_balance
		data.PasswordCorrect = true
	} else {
		data.PasswordWrong = true
	}
	templ.Execute(w, data)

}
