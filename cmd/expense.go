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
