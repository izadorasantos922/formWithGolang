package main

import (
	"log"
	"net/http"
	"webform/handlers"
)

func main() {
	http.HandleFunc("/", handlers.SubscriptionHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
