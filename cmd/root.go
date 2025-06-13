/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// add Command
var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		desc, err := cmd.Flags().GetString("description")
		if err != nil {
			fmt.Println("failed to get description")
		}

		amount, err := cmd.Flags().GetString("amount")
		if err != nil {
			fmt.Println("failed to get amount")
		}

		if desc != "" && amount != "" {
			AddExpense(desc, amount)
		}
	},
}

// list command
var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		listExpense()
	},
}

// summary command
var summaryCmd = &cobra.Command{
	Use: "summary",
	Run: func(cmd *cobra.Command, args []string) {
		summaryExpense()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().String("description", "", "Your expense description")
	rootCmd.PersistentFlags().String("amount", "", "Your expense amount")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(summaryCmd)
}
