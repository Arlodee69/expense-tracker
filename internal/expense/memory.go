package expense

import (
	"sync"
)

type SyncExpense struct {
	mut      sync.Mutex
	expenses []Expense
}

func (e *SyncExpense) Add(value Expense) {
	e.mut.Lock()
	defer e.mut.Unlock()

	e.expenses = append(e.expenses, value)
}

func (e *SyncExpense) List() []Expense {
	e.mut.Lock()
	defer e.mut.Unlock()

	result := make([]Expense, len(e.expenses))
	copy(result, e.expenses)
	return result
}
