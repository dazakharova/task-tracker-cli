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

	newTask := Task{ID: newID, Description: description, CreatedAt: time.Now()}
	tasks = append(tasks, newTask)

	err = Save(file, tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
	return err
}
