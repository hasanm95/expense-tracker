/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"expense-tracker/cmd"
	"expense-tracker/storage"
)

func main() {
	storage.Load()
	cmd.Execute()
}
