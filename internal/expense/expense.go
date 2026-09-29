package expense

import (
	"time"
)
type Expense struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}

