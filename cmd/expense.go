package cmd

import (
	"expense-tracker/storage"
	"expense-tracker/types"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func AddExpense(desc string, amount string) {
	id := int(uuid.New().ID())
	expense := types.Expense{
		ID:          id,
		Description: desc,
		Amount:      "$" + amount,
		CreatedAt:   time.Now(),
	}

	types.AllExpense = append(types.AllExpense, expense)

	err := storage.Storage()

	if err != nil {
		fmt.Println("Failed to add expense")
	}

	fmt.Printf("Expense added successfully (ID: %v)\n", id)
}
