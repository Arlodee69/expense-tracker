package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/Arlodee69/expense-tracker/internal/expense"
)

func HandlerExpenses(w http.ResponseWriter, r *http.Request, store *expense.SyncExpense) {
	if r.Method == http.MethodGet {

		expenses := store.List()
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(expenses)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
	} else if r.Method == http.MethodPost {

		var newExpense expense.Expense
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&newExpense)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		if newExpense.ID < 0 || newExpense.CategoryID < 0 || newExpense.SpentOn == "" || newExpense.Description == "" || newExpense.Amount <= 0 || newExpense.CreatedAt.IsZero() {
			http.Error(w, "Invalid expense data", http.StatusBadRequest)
			return
		}
		store.Add(newExpense)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newExpense)
	}
}
