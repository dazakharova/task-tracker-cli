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
	if err != nil {
		return err
	}
	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
	return nil
}

func ListTasks(file string, status string) error {
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

	filtered := tasks
	if status != "" {
		filtered = nil
		for _, task := range tasks {
			if task.Status == status {
				filtered = append(filtered, task)
			}
		}

		if len(filtered) == 0 {
			fmt.Printf("No tasks with status %q found.\n", status)
			return nil
		}
	}

	fmt.Printf("%-4s %-12s %-17s %s\n", "ID", "Status", "Created", "Description")

	for _, task := range filtered {
		fmt.Printf(
			"%-4d %-12s %-17s %s\n",
			task.ID,
			task.Status,
			task.CreatedAt.Format("2006-01-02 15:04"),
			task.Description,
		)
	}

	return nil
}

func UpdateTask(file string, ID int, description string) error {
	if file == "" {
		return errors.New("filename cannot be empty")
	}

	if description == "" {
		return errors.New("task description is required")
	}

	tasks, err := Load(file)
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == ID {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()

			err = Save(file, tasks)
			if err != nil {
				return err
			}

			fmt.Printf("Task updated successfully (ID: %d)\n", ID)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", ID)
}

func MarkTaskInProgress(file string, ID int) error {
	return markTaskStatus(file, ID, "in progress")
}

func MarkTaskDone(file string, ID int) error {
	return markTaskStatus(file, ID, "done")

}

func markTaskStatus(file string, ID int, status string) error {
	if file == "" {
		return errors.New("filename cannot be empty")
	}

	tasks, err := Load(file)
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == ID {
			tasks[i].Status = status

			err = Save(file, tasks)
			if err != nil {
				return err
			}

			fmt.Printf("Task updated successfully (ID: %d)\n", ID)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", ID)
}

func DeleteTask(file string, ID int) error {
	if file == "" {
		return errors.New("filename cannot be empty")
	}

	tasks, err := Load(file)
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == ID {
			tasks = append(tasks[:i], tasks[i+1:]...)

			err = Save(file, tasks)
			if err != nil {
				return err
			}

			fmt.Printf("Task deleted successfully (ID: %d)\n", ID)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", ID)
}
