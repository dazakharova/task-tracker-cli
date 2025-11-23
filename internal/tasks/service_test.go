package tasks

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	t.Run("Valid task with description and filename", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "tasks.json")

		err := AddTask(filename, "Buy groceries")
		if err != nil {
			t.Errorf("Failed to add task: %v", err)
		}
	})

	t.Run("Missing filename returns error", func(t *testing.T) {
		err := AddTask("", "Buy groceries")

		if err == nil || err.Error() != "filename cannot be empty" {
			t.Fatalf("Expected error %q, got %v", "filename cannot be empty", err)
		}
	})

	t.Run("Missing description returns error", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "tasks.json")

		err := AddTask(filename, "")

		if err == nil || err.Error() != "task description is required" {
			t.Fatalf("Expected error %q, got %v", "task description is required", err)
		}
	})
}

func TestListTasks(t *testing.T) {
	t.Run("Prints tasks for non-empty file (no status filter)", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "tasks.json")

		jsonTasks := `[
		  {
			"id": 1,
			"description": "Buy groceries",
			"status": "todo",
			"created_at": "2025-01-12T15:04:05Z"
		  },
		  {
			"id": 2,
			"description": "Cook dinner",
			"status": "todo",
			"created_at": "2025-01-12T15:04:05Z"
		  }
		]`

		if err := os.WriteFile(filename, []byte(jsonTasks), 0o644); err != nil {
			t.Fatalf("Failed to write temp file: %v", err)
		}

		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		err := ListTasks(filename, "")
		if err != nil {
			t.Fatalf("ListTasks returned error: %v", err)
		}

		w.Close()
		os.Stdout = old

		out, _ := io.ReadAll(r)
		output := string(out)

		got := strings.TrimSpace(output)

		mustContain := []string{
			"ID",
			"Status",
			"Created",
			"Description",
			"1",
			"2",
			"Buy groceries",
			"Cook dinner",
			"todo",
			"2025-01-12 15:04",
		}

		for _, s := range mustContain {
			if !strings.Contains(got, s) {
				t.Fatalf("expected output to contain %q, got:\n%s", s, got)
			}
		}
	})

	t.Run("Prints message when no tasks found", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "tasks.json")

		if err := os.WriteFile(filename, []byte{}, 0o644); err != nil {
			t.Fatalf("Failed to write temp file: %v", err)
		}

		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		err := ListTasks(filename, "")
		if err != nil {
			t.Fatalf("ListTasks returned error: %v", err)
		}

		w.Close()
		os.Stdout = old

		out, _ := io.ReadAll(r)
		output := string(out)

		got := strings.TrimSpace(output)
		want := "No tasks found."

		if !strings.Contains(got, want) {
			t.Fatalf("Expected success message: %s, got %q", want, got)
		}
	})

	t.Run("Returns error when filename is empty", func(t *testing.T) {
		err := ListTasks("", "")
		if err == nil || err.Error() != "filename cannot be empty" {
			t.Fatalf("Expected error %q, got %v", "filename cannot be empty", err)
		}
	})

	t.Run("Prints only tasks with given status", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "tasks.json")

		jsonTasks := `[
		  {
			"id": 1,
			"description": "Buy groceries",
			"status": "todo",
			"created_at": "2025-01-12T15:04:05Z"
		  },
		  {
			"id": 2,
			"description": "Cook dinner",
			"status": "done",
			"created_at": "2025-01-12T16:00:00Z"
		  },
		  {
			"id": 3,
			"description": "Clean kitchen",
			"status": "todo",
			"created_at": "2025-01-12T17:00:00Z"
		  }
		]`

		if err := os.WriteFile(filename, []byte(jsonTasks), 0o644); err != nil {
			t.Fatalf("Failed to write temp file: %v", err)
		}

		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		err := ListTasks(filename, "done")
		if err != nil {
			t.Fatalf("ListTasks returned error: %v", err)
		}

		w.Close()
		os.Stdout = old

		out, _ := io.ReadAll(r)
		output := string(out)
		got := strings.TrimSpace(output)

		if !strings.Contains(got, "Cook dinner") {
			t.Fatalf("expected output to contain %q, got:\n%s", "Cook dinner", got)
		}
		if strings.Contains(got, "Buy groceries") {
			t.Fatalf("did not expect output to contain %q when filtering by status %q, got:\n%s", "Buy groceries", "done", got)
		}
		if strings.Contains(got, "Clean kitchen") {
			t.Fatalf("did not expect output to contain %q when filtering by status %q, got:\n%s", "Clean kitchen", "done", got)
		}
	})

	t.Run("Prints message when no tasks with given status found", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "tasks.json")

		jsonTasks := `[
		  {
			"id": 1,
			"description": "Buy groceries",
			"status": "todo",
			"created_at": "2025-01-12T15:04:05Z"
		  }
		]`

		if err := os.WriteFile(filename, []byte(jsonTasks), 0o644); err != nil {
			t.Fatalf("Failed to write temp file: %v", err)
		}

		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		err := ListTasks(filename, "done")
		if err != nil {
			t.Fatalf("ListTasks returned error: %v", err)
		}

		w.Close()
		os.Stdout = old

		out, _ := io.ReadAll(r)
		output := string(out)
		got := strings.TrimSpace(output)

		want := `No tasks with status "done" found.`

		if !strings.Contains(got, want) {
			t.Fatalf("Expected message %q, got:\n%s", want, got)
		}
	})
}

func createTempTasksFile(t *testing.T, tasks []Task) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")

	if err := Save(path, tasks); err != nil {
		t.Fatalf("failed to save initial tasks: %v", err)
	}

	return path
}

func TestUpdateTask(t *testing.T) {
	t.Run("Updates existing task successfully", func(t *testing.T) {
		now := time.Now()

		initialTasks := []Task{
			{
				ID:          1,
				Description: "First task",
				Status:      "todo",
				CreatedAt:   now.Add(-2 * time.Hour),
			},
			{
				ID:          2,
				Description: "Second task",
				Status:      "in progress",
				CreatedAt:   now.Add(-1 * time.Hour),
			},
		}

		filename := createTempTasksFile(t, initialTasks)

		newDescription := "Updated second task description"

		err := UpdateTask(filename, 2, newDescription)
		if err != nil {
			t.Fatalf("UpdateTask returned error: %v", err)
		}

		updatedTasks, err := Load(filename)
		if err != nil {
			t.Fatalf("Load returned error after update: %v", err)
		}

		if len(updatedTasks) != 2 {
			t.Fatalf("Expected 2 tasks, got %d", len(updatedTasks))
		}

		var updated Task
		for _, task := range updatedTasks {
			if task.ID == 2 {
				updated = task
				break
			}
		}

		if updated.ID == 0 {
			t.Fatalf("Task with ID 2 not found after update")
		}

		if updated.Description != newDescription {
			t.Errorf("Expected description %q, got %q", newDescription, updated.Description)
		}

		if updated.UpdatedAt.IsZero() {
			t.Errorf("Expected UpdatedAt to be set, but it is zero")
		}

		for _, task := range updatedTasks {
			if task.ID == 1 && task.Description != "First task" {
				t.Errorf("Task with ID 1 should not be modified, got description %q", task.Description)
			}
		}
	})

	t.Run("Returns error when task not found", func(t *testing.T) {
		initialTasks := []Task{
			{
				ID:          1,
				Description: "First task",
				Status:      "todo",
				CreatedAt:   time.Now(),
			},
		}

		filename := createTempTasksFile(t, initialTasks)

		err := UpdateTask(filename, 99, "Does not matter")

		if !strings.Contains(err.Error(), "not found") {
			t.Fatalf("Expected error to mention %q, got %q", "not found", err.Error())
		}
	})

	t.Run("Returns error when description is empty", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "tasks.json")

		err := UpdateTask(filename, 1, "")

		if err == nil || err.Error() != "task description is required" {
			t.Fatalf("Expected error %q, got %v", "task description is required", err)
		}
	})

	t.Run("Returns error when filename is empty", func(t *testing.T) {
		err := UpdateTask("", 1, "Some description")

		if err == nil || err.Error() != "filename cannot be empty" {
			t.Fatalf("Expected error %q, got %v", "filename cannot be empty", err)
		}
	})
}
