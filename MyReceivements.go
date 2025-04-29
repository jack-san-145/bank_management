// package main

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// type receivecardType struct {
// 	Present          bool
// 	Sender_id        int
// 	Receiver_id      int
// 	Amount           int
// 	Transaction_date string
// 	Sender_name      string
// 	Receiver_name    string
// }

// type receivedataCard struct {
// 	Card1  cardType
// 	Card2  cardType
// 	Card3  cardType
// 	Card4  cardType
// 	Card5  cardType
// 	Card6  cardType
// 	Card7  cardType
// 	Card8  cardType
// 	Card9  cardType
// 	Card10 cardType
// }

// func MyreceivementHandler(w http.ResponseWriter, r *http.Request) {
// 	templ, _ := template.ParseFiles("HTML/MyReceivements.html")

// 	data := dataCard{}

// 	account_no, _, _, _, _, _ := FindUserFromCookie(w, r)
// 	query := "select * from Transaction where receiver_id = ? order by transaction_date desc limit 10"
// 	cards := []*cardType{
// 		&data.Card1, &data.Card2, &data.Card3, &data.Card4, &data.Card5,
// 		&data.Card6, &data.Card7, &data.Card8, &data.Card9, &data.Card10,
// 	}
// 	rows, _ := Db.Query(query, account_no)
// 	i := 0
// 	for rows.Next() {
// 		card := cards[i]
// 		card.Present = true
// 		err := rows.Scan(
// 			&card.Sender_id,
// 			&card.Receiver_id,
// 			&card.Amount,
// 			&card.Transaction_date,
// 			&card.Sender_name,
// 			&card.Receiver_name)
// 		if err != nil {
// 			log.Print("error")
// 			continue
// 		}
// 		i++
// 	}
// 	// fmt.Println(data)
// 	// fmt.Println(data.Card1.Sender_name)
// 	err := templ.Execute(w, data)
// 	if err != nil {
// 		log.Println("Template execution error:", err)
// 	}

// }

package main

import (
	"html/template"
	"log"
	"net/http"
)

func MyreceivementHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template safely
	templ, err := template.ParseFiles("HTML/MyReceivements.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		log.Println("Template parsing error:", err)
		return
	}

	// Initialize card data
	data := dataCard{}

	// Get account number from cookie (you may want to validate if it's OK)
	account_no, _, _, _, _, _, err := FindUserFromCookie(w, r)
	if err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		log.Println("Cookie retrieval error:", err)
		return
	}

	// Query the DB for latest 10 received transactions
	query := "SELECT sender_id, receiver_id, amount, transaction_date, sender_name, receiver_name FROM Transaction WHERE receiver_id = ? ORDER BY transaction_date DESC LIMIT 10"
	rows, err := Db.Query(query, account_no)
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		log.Println("DB query error:", err)
		return
	}
	defer rows.Close()

	// Prepare a list of card pointers
	cards := []*cardType{
		&data.Card1, &data.Card2, &data.Card3, &data.Card4, &data.Card5,
		&data.Card6, &data.Card7, &data.Card8, &data.Card9, &data.Card10,
	}

	i := 0
	for rows.Next() {
		if i >= len(cards) {
			break // Prevent index out of range
		}
		card := cards[i]
		err := rows.Scan(
			&card.Sender_id,
			&card.Receiver_id,
			&card.Amount,
			&card.Transaction_date,
			&card.Sender_name,
			&card.Receiver_name,
		)
		if err != nil {
			log.Println("Row scan error:", err)
			continue
		}
		card.Present = true
		i++
	}

	// Execute the template with the filled data
	err = templ.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution failed", http.StatusInternalServerError)
		log.Println("Template execution error:", err)
	}
}
