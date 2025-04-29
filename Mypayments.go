package main

import (
	"html/template"
	"log"
	"net/http"
)

type cardType struct {
	Present          bool
	Sender_id        int
	Receiver_id      int
	Amount           int
	Transaction_date string
	Sender_name      string
	Receiver_name    string
}

type dataCard struct {
	Card1  cardType
	Card2  cardType
	Card3  cardType
	Card4  cardType
	Card5  cardType
	Card6  cardType
	Card7  cardType
	Card8  cardType
	Card9  cardType
	Card10 cardType
}

func MypaymentHandler(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.ParseFiles("HTML/MyPayments.html")

	data := dataCard{}

	account_no, _, _, _, _, _, _ := FindUserFromCookie(w, r)
	query := "select * from Transaction where sender_id = ? order by transaction_date desc limit 10"
	cards := []*cardType{
		&data.Card1, &data.Card2, &data.Card3, &data.Card4, &data.Card5,
		&data.Card6, &data.Card7, &data.Card8, &data.Card9, &data.Card10,
	}
	rows, _ := Db.Query(query, account_no)
	i := 0
	for rows.Next() {
		card := cards[i]
		card.Present = true
		err := rows.Scan(
			&card.Sender_id,
			&card.Receiver_id,
			&card.Amount,
			&card.Transaction_date,
			&card.Sender_name,
			&card.Receiver_name)
		if err != nil {
			log.Print("error")
			continue
		}
		i++
	}
	// fmt.Println(data)
	// fmt.Println(data.Card1.Sender_name)
	err := templ.Execute(w, data)
	if err != nil {
		log.Println("Template execution error:", err)
	}

}
