package types

import "time"

type Expense struct {
	ID          int
	Description string
	Amount      string
	CreatedAt   time.Time
}

var AllExpense []Expense
