package main

import (
	"net/http"

	"overlix-backend/internal/handler"
)

func main() {

	http.HandleFunc("/invite", handler.GenerateInvite)
	http.HandleFunc("/invite/", handler.ResolveInvite)
	http.HandleFunc("/checkin", handler.DailyCheckin)

	http.ListenAndServe(":8080", nil)

}
