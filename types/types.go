package types

import "time"

type Expense struct {
	ID          int
	Description string
	Amount      string
	Date        time.Time
}

var AllExpense []Expense
