package expense

import (
	"time"
)

type Expense struct {
	ID          int64     `json:"id"`
	CategoryID  int64     `json:"category_id"`
	Amount      int64     `json:"amount"`
	SpentOn     string    `json:"spent_on"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
