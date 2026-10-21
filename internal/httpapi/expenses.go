package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/Arlodee69/expense-tracker/internal/expense"
)

func HandlerExpenses(w http.ResponseWriter, r *http.Request, store *expense.PostgresPool) {
	if r.Method == http.MethodGet {

		expenses, err := store.List(r.Context())
		if err != nil {
			log.Printf("list expenses: %v", err)
			http.Error(w, "Failed to retrieve expenses", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(expenses)
		if err != nil {
			log.Printf("encode expenses response: %v", err)
			return
		}
	} else if r.Method == http.MethodPost {

		var newExpense expense.Expense
		err := json.NewDecoder(r.Body).Decode(&newExpense)
		if err != nil {
			http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
			return
		}
		if newExpense.ID < 0 || newExpense.CategoryID <= 0 || newExpense.Description == "" || newExpense.Amount <= 0 || utf8.RuneCountInString(newExpense.Description) > 1000 {
			http.Error(w, "Invalid expense fields", http.StatusBadRequest)
			return
		}
		if _, err := time.Parse(time.DateOnly, newExpense.SpentOn); err != nil {
			http.Error(w, "spent_on must be a valid date in YYYY-MM-DD format", http.StatusBadRequest)
			return
		}
		created, err := store.Add(r.Context(), newExpense)
		if err != nil {
			log.Printf("add expense: %v", err)
			http.Error(w, "Failed to save expense", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(created)
		if err != nil {
			log.Printf("encode created expense response: %v", err)
			return
		}
	} else if r.Method == http.MethodDelete {
		deleteID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil || deleteID <= 0 {
			http.Error(w, "Expense id must be a positive integer", http.StatusBadRequest)
			return
		}
		deleted, err := store.Delete(r.Context(), deleteID)
		if err != nil {
			log.Printf("delete expense %d: %v", deleteID, err)
			http.Error(w, "Failed to delete expense", http.StatusInternalServerError)
			return
		}
		if !deleted {
			http.Error(w, "Expense not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	} else {
		log.Printf("method %v not allowed", r.Method)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
