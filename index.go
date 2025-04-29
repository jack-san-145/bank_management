package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	CurrentUserName     string
	CurrentUserPassword string
	CurrentUserMobile   string
	CurrentUserAddress  string
	CurrentUserBalance  int
)

func ServeIndexPage(w http.ResponseWriter, r *http.Request) {
	// temp, err := template.ParseFiles("HTML/Home.html")
	// if err != nil {
	// 	http.Error(w, "unable to load the Home page", http.StatusInternalServerError)
	// 	return
	// }
	temp := LoadhtmlPage(w, "index.html")
	temp.Execute(w, nil)

}

func FindCurrentUser() {
	var err error
	query_for_details := "select cus_name,password,address,mobile_no from CustomersList where acc_no = ?"
	err = Db.QueryRow(query_for_details, CurrentUserAc).Scan(&CurrentUserName, &CurrentUserPassword, &CurrentUserAddress, &CurrentUserMobile)
	query_for_balance := "Select balance from CustomerBalance where acc_no = ?"
	err = Db.QueryRow(query_for_balance, CurrentUserAc).Scan(&CurrentUserBalance)
	fmt.Print(err)
}

func FindUserFromCookie(w http.ResponseWriter, r *http.Request) (int, string, string, string, string, int, error) {
	var err error
	var (
		account_no       int
		account_name     string
		account_password string
		account_address  string
		account_mobile   string
		account_balance  int
	)
	cookie, err := r.Cookie("account_number")
	if err != nil {
		return 0, "", "", "", "", 0, fmt.Errorf("Error")
	}
	if cookie == nil {
		log.Println("sunni")
	}
	UserAc := cookie.Value
	query_for_details := "select * from CustomersList where acc_no = ?"
	err = Db.QueryRow(query_for_details, UserAc).Scan(&account_no, &account_name, &account_password, &account_address, &account_mobile)

	query_for_balance := "Select balance from CustomerBalance where acc_no = ?"
	err = Db.QueryRow(query_for_balance, UserAc).Scan(&account_balance)

	fmt.Println(account_no, account_name, account_password, account_address, account_mobile, account_balance)

	return account_no, account_name, account_password, account_address, account_mobile, account_balance, nil

}
