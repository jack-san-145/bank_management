package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var Db *sql.DB

func main() {

	http.HandleFunc("/bank/serveRegister", ServeRegistrationPage)
	http.HandleFunc("/bank/register", RegistrationHandler)
	// http.HandleFunc("/password", passwordHandler)
	http.HandleFunc("/bank/login", LoginHandler)
	http.HandleFunc("/bank/serveIndex", ServeIndexPage)
	http.HandleFunc("/asserts/", LoadAsserts)
	http.HandleFunc("/bank/main", ServeLoginPage)
	http.HandleFunc("/bank/HTML/", HtmlHandler)
	http.HandleFunc("/bank/Details", DetailsHandler)
	http.HandleFunc("/bank/AccountHolder", AccountHolderHandler)
	http.HandleFunc("/bank/serveTransaction", ServeTransactionPage)
	http.HandleFunc("/bank/Transaction", TransactionHandler)
	http.HandleFunc("/bank/serveBalancePage", ServeBalancePage)
	http.HandleFunc("/bank/Balance", BalanceHandler)
	http.HandleFunc("/bank/Mypayment", MypaymentHandler)
	http.HandleFunc("/bank/Myreceivement", MyreceivementHandler)
	http.HandleFunc("/bank/logout", LogoutHandler)
	fmt.Println("Server is running on the Port 8989")

	var err error

	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3307)/Bank")
	if err != nil {
		log.Fatal("Database Connection failed - ", err)
	} else {
		fmt.Println("Database Connected")
	}
	defer Db.Close()

	err = http.ListenAndServe(":8989", nil)
	if err != nil {
		fmt.Println("Server failure")
	}

	// func enableCORS(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		// Set CORS headers
	// 		w.Header().Set("Access-Control-Allow-Origin", "*") // Or restrict to specific domain
	// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// 		// Handle preflight requests
	// 		if r.Method == "OPTIONS" {
	// 			w.WriteHeader(http.StatusOK)
	// 			return
	// 		}

	// 		next.ServeHTTP(w, r)
	// 	})
	// }

	// func main() {

	// 	http.Handle("/bank/serveRegister", enableCORS(http.HandlerFunc(ServeRegistrationPage)))
	// 	http.Handle("/bank/register", enableCORS(http.HandlerFunc(RegistrationHandler)))
	// 	http.Handle("/bank/login", enableCORS(http.HandlerFunc(LoginHandler)))
	// 	http.Handle("/bank/serveIndex", enableCORS(http.HandlerFunc(ServeIndexPage)))
	// 	http.Handle("/asserts/", enableCORS(http.HandlerFunc(LoadAsserts)))
	// 	http.Handle("/bank/main", enableCORS(http.HandlerFunc(ServeLoginPage)))
	// 	http.Handle("/bank/HTML/", enableCORS(http.HandlerFunc(HtmlHandler)))
	// 	http.Handle("/bank/Details", enableCORS(http.HandlerFunc(DetailsHandler)))
	// 	http.Handle("/bank/AccountHolder", enableCORS(http.HandlerFunc(AccountHolderHandler)))
	// 	http.Handle("/bank/serveTransaction", enableCORS(http.HandlerFunc(ServeTransactionPage)))
	// 	http.Handle("/bank/Transaction", enableCORS(http.HandlerFunc(TransactionHandler)))
	// 	http.Handle("/bank/serveBalancePage", enableCORS(http.HandlerFunc(ServeBalancePage)))
	// 	http.Handle("/bank/Balance", enableCORS(http.HandlerFunc(BalanceHandler)))
	// 	http.Handle("/bank/Mypayment", enableCORS(http.HandlerFunc(MypaymentHandler)))
	// 	http.Handle("/bank/Myreceivement", enableCORS(http.HandlerFunc(MyreceivementHandler)))

	// 	fmt.Println("Server is running on the Port 8989")

	// 	var err error

	// 	Db, err = sql.Open("mysql", "root@tcp(localhost:3306)/Bank")
	// 	if err != nil {
	// 		log.Fatal("Database Connection failed - ", err)
	// 	} else {
	// 		fmt.Println("Database Connected")
	// 	}
	// 	defer Db.Close()

	// 	err = http.ListenAndServe("0.0.0.0:8989", nil)

	// 	if err != nil {
	// 		fmt.Println("Server failure")
	// 	}
	// }

}
