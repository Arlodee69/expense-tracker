package main

import (
	"log"
	"net/http"

	"github.com/Arlodee69/expense-tracker/internal/expense"
	"github.com/Arlodee69/expense-tracker/internal/httpapi"
)

func main() {

	storage := &expense.SyncExpense{}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /expenses", func(w http.ResponseWriter, r *http.Request) {
		httpapi.HandlerExpenses(w, r, storage)
	})
	mux.HandleFunc("POST /expenses", func(w http.ResponseWriter, r *http.Request) {
		httpapi.HandlerExpenses(w, r, storage)
	})
	log.Println("server started: http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
