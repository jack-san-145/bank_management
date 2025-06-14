package main

import (
	"html/template"
	"net/http"
)

type accountInfo struct {
	AccountNumber int
	Name          string
	MobileNumber  string
	Address       string
}

func AccountHolderHandler(w http.ResponseWriter, r *http.Request) {

	if !checkForCookie(w, r) {
		return
	}

	var (
		account_no      int
		account_name    string
		account_address string
		account_mobile  string
	)
	account_no, account_name, _, account_address, account_mobile, _, _ = FindUserFromCookie(w, r)

	templ, _ := template.ParseFiles("HTML/AccountHolder.html")
	data := accountInfo{
		AccountNumber: account_no,
		Name:          account_name,
		MobileNumber:  account_mobile,
		Address:       account_address,
	}
	templ.Execute(w, data)

}
