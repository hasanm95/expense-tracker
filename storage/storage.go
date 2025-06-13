package storage

import (
	"encoding/json"
	"expense-tracker/types"
	"fmt"
	"os"
)

func Storage() error {
	file, err := os.Create("expenses.json")

	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(types.AllExpense, "", "   ")

	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if _, err := file.Write(jsonData); err != nil {
		return fmt.Errorf("failed to write data in json: %w", err)
	}

	return nil
}

func Load() error {
	fileBytes, err := os.ReadFile("expenses.json")

	if err != nil {
		if os.IsNotExist(err) {
			types.AllExpense = []types.Expense{}
			return nil
		}
		return fmt.Errorf("failed to read file: %w", err)
	}

	if err = json.Unmarshal(fileBytes, &types.AllExpense); err != nil {
		return fmt.Errorf("failed to parse expense data: %w", err)
	}

	return nil
}
