package cmd

import (
	"expense-tracker/storage"
	"expense-tracker/types"
	"fmt"
	"os"
	"regexp"
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

func summaryExpense() {
	var sum int = 0
	for _, exp := range types.AllExpense {
		re := regexp.MustCompile(`[^0-9.]`)
		amount := re.ReplaceAllString(exp.Amount, "")
		val, err := strconv.Atoi(amount)
		if err != nil {
			fmt.Printf("Falied to conver amount: %v\n", err)
		}
		sum += val
	}
	fmt.Printf("Total Expenses: $%v\n", sum)
}

func deleteExpense(id int) {
	var indexToRemove = -1
	for i, exp := range types.AllExpense {
		if id == exp.ID {
			indexToRemove = i
			break
		}
	}

	if indexToRemove == -1 {
		fmt.Printf("Expense with ID %d not found", id)
	}
	types.AllExpense = append(types.AllExpense[:indexToRemove], types.AllExpense[indexToRemove+1:]...)

	if err := storage.Storage(); err != nil {
		fmt.Printf("Failed to store %v", err)
	}

	fmt.Printf("Task removed successfully (ID: %d)\n", id)
}
