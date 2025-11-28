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
			exitUsageError("Error: missing task description.")
		}

		taskText := strings.Join(args[1:], " ")

		if err := tasks.AddTask(file, taskText); err != nil {
			exitFatalError("Error adding task: %v\n", err)
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
				exitUsageError("Error: invalid task status.\nAllowed statuses: todo, in progress, done.")
			}
		} else {
			status = ""
		}

		if err := tasks.ListTasks(file, status); err != nil {
			exitFatalError("Error listing tasks: %v\n", err)
		}
	case "update":
		if len(args) < 2 {
			exitUsageError("Error: missing task ID and description.")
		} else if len(args) < 3 {
			exitUsageError("Error: missing task description.")
		}

		idStr := args[1]
		taskID, err := strconv.Atoi(idStr)
		if err != nil {
			exitFatalError("Error: invalid task ID: %v\n", err)
		}

		newDescription := strings.Join(args[2:], " ")

		if err := tasks.UpdateTask(file, taskID, newDescription); err != nil {
			exitFatalError("Error updating task: %v\n", err)
		}
	case "mark-in-progress":
		if len(args) < 2 {
			exitUsageError("Error: missing task ID.")
		}

		idStr := args[1]

		err := handleMarkStatus("in progress", file, idStr)
		if err != nil {
			exitFatalError("Error marking task 'in progress'", err)
		}
	case "mark-done":
		if len(args) < 2 {
			exitUsageError("Error: missing task ID.")
		}

		idStr := args[1]

		err := handleMarkStatus("done", file, idStr)
		if err != nil {
			exitFatalError("Error marking task 'done'", err)
		}
	case "delete":
		if len(args) < 2 {
			exitUsageError("Error: missing task ID.")
		}

		idStr := args[1]
		taskID, err := strconv.Atoi(idStr)
		if err != nil {
			exitFatalError("Error: invalid task ID: %v\n", err)
		}

		err = tasks.DeleteTask(file, taskID)
		if err != nil {
			exitFatalError("Error deleting task: %v\n", err)
		}
	default:
		exitUsageError("Invalid command: " + command)
	}
}

func exitUsageError(message string) {
	fmt.Fprintln(os.Stderr, message)
	fmt.Fprintln(os.Stderr)
	showHelp()
	os.Exit(1)
}

func exitFatalError(message string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
	os.Exit(1)
}

func handleMarkStatus(status string, filename string, idStr string) error {
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid task ID %q: %w", idStr, err)
	}

	switch status {
	case "in progress":
		return tasks.MarkTaskInProgress(filename, taskID)
	case "done":
		return tasks.MarkTaskDone(filename, taskID)
	default:
		return fmt.Errorf("unsupported status %q", status)
	}
}
