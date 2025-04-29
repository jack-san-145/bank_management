package main

import (
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type userDetalils struct {
	AccountNumber int
	Name          string
	MobileNumber  string
	Address       string
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	var account_no int
	query := "Select acc_no from CustomersList where password = ? and cus_name=? and mobile_no=?"
	_ = Db.QueryRow(query, Customer_password, Customer_name, Customer_mobile).Scan(&account_no)
	templ, _ := template.ParseFiles("HTML/Details.html")
	data := userDetalils{
		AccountNumber: account_no,
		Name:          Customer_name,
		MobileNumber:  Customer_mobile,
		Address:       Customer_address,
	}
	templ.Execute(w, data)
}
