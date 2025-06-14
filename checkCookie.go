package main

import (
	"fmt"
	"net/http"
)

func checkForCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		fmt.Println("Cookie found - ", cookie.Value)
		return true
	} else {
		ServeLoginPage(w, r)
		return false
	}
}
