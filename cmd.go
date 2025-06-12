package main

import (
	"flag"
	"fmt"
)

func cmdFlags() error {
	// Parse flags
	flag.Parse()

	// Get remaining arguments after flags
	args := flag.Args()

	// Handle commands
	if err := handleCommands(args); err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func handleCommands(args []string) error {
	// Check if any command was provided
	if len(args) == 0 && !flag.Parsed() {
		return fmt.Errorf("no command provided")
	}

	// Handle arguments
	if len(args) > 0 {
		command := args[0]
		fmt.Println("Command", command)
		switch command {
		case "add":
			fmt.Println("Add command", args)
		default:
			return fmt.Errorf("unknown command: %s", command)
		}
	}

	return nil
}
