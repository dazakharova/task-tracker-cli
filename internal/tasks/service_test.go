package tasks

import (
	"path/filepath"
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
