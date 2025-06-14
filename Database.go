package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func mainDatabase() {
	account_number := 12
	amount := 2000
	query := "update CustomerBalance set balance = ? where acc_no = ? "
	result, err := Db.Exec(query, amount, account_number)
	if err != nil {
		fmt.Println("Error occured")
		return
	} else {
		fmt.Println("Result = ", result)
	}
}

func FindCurrentBalance(id int) int {
	var balance int
	query := "select balance from CustomerBalance where acc_no= ? "
	err := Db.QueryRow(query, id).Scan(&balance)
	i := 0
	if err != nil {
		i++
	}
	return balance
}

func UpdateBalance(amount int, account int) {
	query := "update CustomerBalance set balance = ? where acc_no = ? "
	result, _ := Db.Exec(query, amount, account)
	fmt.Println("result = ", result)
}

func TransactionRecord(senderId int, receiverId int, senderName string, receiverName string, amount int) {
	query := "insert into Transaction(sender_id,receiver_id,sender_name,receiver_name,amount) values(?,?,?,?,?)"
	result, _ := Db.Exec(query, senderId, receiverId, senderName, receiverName, amount)
	fmt.Println("Transaction record - ", result)
}

func DeleteSession(session_id string) {
	query := "delete from Sessions where session_id = ? "
	_, err := Db.Exec(query, session_id)
	if err != nil {
		fmt.Println("error while deleting session ")
	}
}
