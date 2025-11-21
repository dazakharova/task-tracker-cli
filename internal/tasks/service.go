package tasks

import (
	"errors"
	"fmt"
	"time"
)

func AddTask(file, description string) error {
	if description == "" {
		return errors.New("task description is required")
	}

	if file == "" {
		return errors.New("filename cannot be empty")
	}

	tasks, err := Load(file)
	if err != nil {
		return err
	}

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{ID: newID, Description: description, Status: "todo", CreatedAt: time.Now()}
	tasks = append(tasks, newTask)

	err = Save(file, tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
	return err
}

func ListTasks(file string) error {
	if file == "" {
		return errors.New("filename cannot be empty")
	}

	tasks, err := Load(file)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	for _, task := range tasks {
		fmt.Printf("#%d [%s] %s %s\n",
			task.ID,
			task.CreatedAt.Format("2006-01-02 15:04"),
			task.Description,
			task.Status,
		)
	}
	return nil
}
