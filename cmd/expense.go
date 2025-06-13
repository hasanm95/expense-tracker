package cmd

import (
	"expense-tracker/storage"
	"expense-tracker/types"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/google/uuid"
)

func AddExpense(desc string, amount string) {
	id := int(uuid.New().ID())
	expense := types.Expense{
		ID:          id,
		Description: desc,
		Amount:      "$" + amount,
		Date:        time.Now(),
	}

	types.AllExpense = append(types.AllExpense, expense)

	err := storage.Storage()

	if err != nil {
		fmt.Println("Failed to add expense")
	}

	fmt.Printf("Expense added successfully (ID: %v)\n", id)
}

func listExpense() {
	t := table.New(os.Stdout)
	t.AddHeaders("ID", "Date", "Description", "Amount")

	for _, exp := range types.AllExpense {
		t.AddRow(strconv.Itoa(exp.ID), exp.Date.Format(time.RFC1123), exp.Description, exp.Amount)
	}

	t.Render()
}
