package main

import (
	"fmt"
	"net/http"
)

type Transaction struct {
	Title  string
	Amount int64
	Type   int
}

func main() {
	http.HandleFunc("/api/welcome", welcomeHandler)
	http.HandleFunc("/", status)
	http.ListenAndServe(":8080", nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to finance system")
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Finance System")
}
