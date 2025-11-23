package main

import (
	"TaskTrackerCLI/internal/tasks"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func showHelp() {
	fmt.Println("Usage:")
	fmt.Println("  task-cli add <task description>")
	fmt.Println("  task-cli list [status]")
	fmt.Println("  task-cli update <id> <new description>")
	fmt.Println("  task-cli mark-in-progress <id>")
	fmt.Println("  task-cli mark-done <id>")
	fmt.Println("  task-cli delete <id>")
	fmt.Println()
	fmt.Println("Status values for list:")
	fmt.Println("  todo")
	fmt.Println("  in progress")
	fmt.Println("  done")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println(`  task-cli add "Buy groceries"`)
	fmt.Println(`  task-cli list`)
	fmt.Println(`  task-cli list done`)
	fmt.Println(`  task-cli list "in progress"`)
	fmt.Println(`  task-cli update 1 "Buy groceries and cook dinner"`)
	fmt.Println(`  task-cli mark-in-progress 3`)
	fmt.Println(`  task-cli mark-done 1`)
	fmt.Println(`  task-cli delete 2`)
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
		var validStatuses = map[string]bool{
			"todo":        true,
			"in progress": true,
			"done":        true,
		}

		var status string
		if len(args) > 1 {
			status = strings.Join(args[1:], " ")

			if !validStatuses[status] {
				fmt.Printf("Error: invalid status: %s.", status)
				fmt.Println("Allowed statuses: todo, in progress, done.")
				fmt.Println()
				showHelp()
				os.Exit(1)
			}
		} else {
			status = ""
		}

		if err := tasks.ListTasks(file, status); err != nil {
			fmt.Fprintf(os.Stderr, "Error listing tasks: %v\n", err)
			os.Exit(1)
		}
	case "update":
		if len(args) < 2 {
			fmt.Println("Error: missing task ID and description.")
			fmt.Println()
			showHelp()
			os.Exit(1)
		} else if len(args) < 3 {
			fmt.Println("Error: missing task description.")
			fmt.Println()
			showHelp()
			os.Exit(1)
		}

		idStr := args[1]
		taskID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: invalid task ID %q: %v\n", idStr, err)
			os.Exit(1)
		}

		newDescription := strings.Join(args[2:], " ")

		if err := tasks.UpdateTask(file, taskID, newDescription); err != nil {
			fmt.Fprintf(os.Stderr, "Error updating task: %v\n", err)
			os.Exit(1)
		}
	case "mark-in-progress":
		if len(args) < 2 {
			fmt.Println("Error: missing task ID.")
			fmt.Println()
			showHelp()
			os.Exit(1)
		}

		idStr := args[1]
		taskID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: invalid task ID %q: %v\n", idStr, err)
			os.Exit(1)
		}

		err = tasks.MarkTaskInProgress(file, taskID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marking task 'in progress': %v\n", err)
			os.Exit(1)
		}
	case "mark-done":
		if len(args) < 2 {
			fmt.Println("Error: missing task ID.")
			fmt.Println()
			showHelp()
			os.Exit(1)
		}

		idStr := args[1]
		taskID, err := strconv.Atoi(idStr)

		err = tasks.MarkTaskDone(file, taskID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marking task 'done': %v\n", err)
			os.Exit(1)
		}
	case "delete":
		if len(args) < 2 {
			fmt.Println("Error: missing task ID.")
			fmt.Println()
			showHelp()
			os.Exit(1)
		}

		idStr := args[1]
		taskID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: invalid task ID %q: %v\n", idStr, err)
			os.Exit(1)
		}

		err = tasks.DeleteTask(file, taskID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting task: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Invalid command: %s\n\n", command)
		showHelp()
		os.Exit(1)
	}
}
