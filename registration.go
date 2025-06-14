package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Customer_name     string
	Customer_password string
	Customer_mobile   string
	Customer_address  string
)

func ServeRegistrationPage(w http.ResponseWriter, r *http.Request) {

	temp := LoadhtmlPage(w, "Registration.html")
	temp.Execute(w, nil)

}

func inserting(customer_name, customer_password, customer_mobile, customer_address string) {

	query := "insert into CustomersList(cus_name,password,address,mobile_no) values(?,?,?,?)"
	_, err := Db.Exec(query, customer_name, customer_password, customer_address, customer_mobile)
	if err != nil {
		fmt.Println("something error while inserting customer details - ", err)
		return
	}

	var account_no int
	query = "Select acc_no from CustomersList where password = ? and cus_name=?"
	_ = Db.QueryRow(query, customer_password, customer_name).Scan(&account_no)

	query = "insert into CustomerBalance(acc_no,balance) values(?,?)"
	_, err = Db.Exec(query, account_no, 1000)
	if err != nil {
		fmt.Println("something error while Adding balance")
		return
	}

}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "bank/register", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	Customer_name = r.FormValue("Customer_name")
	Customer_password = r.FormValue("Customer_password")
	Customer_mobile = r.FormValue("mobile")
	Customer_address = r.FormValue("Customer_address")

	var (
		ac      int
		name    string
		pass    string
		address string
		mobile  string
	)
	type pageData struct {
		ShowMessage bool
	}
	query := "select * from CustomersList where mobile_no in(?)"
	result := Db.QueryRow(query, Customer_mobile)
	err := result.Scan(&ac, &name, &pass, &address, &mobile)
	if err == sql.ErrNoRows {
		inserting(Customer_name, Customer_password, Customer_mobile, Customer_address)
		http.Redirect(w, r, "/bank/Details", http.StatusSeeOther)

	} else {
		if name == Customer_name {
			templ, _ := template.ParseFiles("HTML/Registration.html")
			data := pageData{ShowMessage: true}
			templ.Execute(w, data)
		} else {
			inserting(Customer_name, Customer_password, Customer_mobile, Customer_address)
			http.Redirect(w, r, "/bank/Details", http.StatusSeeOther)
		}

	}

}
