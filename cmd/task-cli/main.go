package main

import (
	"TaskTrackerCLI/internal/tasks"
	"fmt"
	"os"
	"strings"
)

func showHelp() {
	fmt.Println("Usage:")
	fmt.Println("  task-cli add <task description>")
	fmt.Println("  task-cli list")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println(`  task-cli add "Buy groceries"`)
	fmt.Println(`  task-cli list`)
}

func main() {
	file := "tasks.json"
	args := os.Args[1:]

	if len(args) == 0 {
		showHelp()
		return
	}

	if args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		showHelp()
		return
	}

	command := args[0]

	switch command {
	case "add":
		if len(args) < 2 {
			fmt.Println("Error: missing task description.")
			fmt.Println()
			showHelp()
			os.Exit(1)
		}

		taskText := strings.Join(args[1:], " ")

		if err := tasks.AddTask(file, taskText); err != nil {
			fmt.Fprintf(os.Stderr, "Error adding task: %v\n", err)
			os.Exit(1)
		}
	case "list":
		if err := tasks.ListTasks(file); err != nil {
			fmt.Fprintf(os.Stderr, "Error listing tasks: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Invalid command: %s\n\n", command)
		os.Exit(1)
	}
}
