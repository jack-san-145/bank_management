package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type transactionStatus struct {
	Success   bool
	Invalid   bool
	Nobalance bool
	Password  bool
}

func ServeTransactionPage(w http.ResponseWriter, r *http.Request) {

	temp := LoadhtmlPage(w, "Transaction.html")
	temp.Execute(w, nil)

}

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	var (
		account_no       int
		account_name     string
		account_password string
		account_balance  int
	)
	templ, _ := template.ParseFiles("HTML/Transaction.html")
	data := transactionStatus{false, false, false, false}

	var (
		r_name string
	)
	r.ParseForm()
	receiver_ac := r.FormValue("acc-number ")
	receiver_name := r.FormValue("acc-holder")
	amount_to_pay := r.FormValue("amount")
	user_password := r.FormValue("password")
	amount_to_pay_int, _ := strconv.Atoi(amount_to_pay)
	receiver_ac_int, _ := strconv.Atoi(receiver_ac)
	fmt.Println(receiver_ac, "/n", receiver_name, "/n", amount_to_pay, "/n", user_password)
	query := "select cus_name from CustomersList where acc_no = ? "
	err := Db.QueryRow(query, receiver_ac).Scan(&r_name)
	if err == sql.ErrNoRows {
		data.Invalid = true

	} else if receiver_name == r_name {
		account_no, account_name, account_password, _, _, account_balance, _ = FindUserFromCookie(w, r)

		if amount_to_pay_int > account_balance {
			data.Nobalance = true
		} else {
			if user_password == account_password {
				receiver_current_balance := FindCurrentBalance(receiver_ac_int)

				UpdateBalance((receiver_current_balance + amount_to_pay_int), receiver_ac_int)
				UpdateBalance((account_balance - amount_to_pay_int), account_no)
				TransactionRecord(account_no, receiver_ac_int, account_name, receiver_name, amount_to_pay_int)
				data.Success = true
			} else {
				data.Password = true
			}

		}
	} else {
		data.Invalid = true
	}

	err = templ.Execute(w, data)
	if err != nil {
		log.Println("Template execution error:", err)
	}
	fmt.Printf("Success: %v, Invalid: %v, NoBalance: %v\n", data.Success, data.Invalid, data.Nobalance)

}
