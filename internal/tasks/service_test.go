package tasks

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
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
	t.Run("Prints tasks for non-empty file", func(t *testing.T) {
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

		err := ListTasks(filename)
		if err != nil {
			t.Fatalf("ListTasks returned error: %v", err)
		}

		w.Close()
		os.Stdout = old

		out, _ := io.ReadAll(r)
		output := string(out)

		got := strings.TrimSpace(output)

		mustContain := []string{
			"ID", // header present
			"Status",
			"Created",
			"Description",
			"1", // IDs present
			"2",
			"Buy groceries", // descriptions present
			"Cook dinner",
			"todo", // status present at least somewhere
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

		err := ListTasks(filename)
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
		err := ListTasks("")
		if err == nil || err.Error() != "filename cannot be empty" {
			t.Fatalf("Expected error %q, got %v", "filename cannot be empty", err)
		}
	})
}
