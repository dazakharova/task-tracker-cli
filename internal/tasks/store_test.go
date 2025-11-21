package tasks

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLoad(t *testing.T) {
	t.Run("Non-empty file with valid tasks", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "tasks.json")

		jsonTasks := `[
		  {
			"id": 1,
			"description": "Buy groceries",
			"created_at": "2025-01-12T15:04:05Z"
		  },
		  {
			"id": 2,
			"description": "Cook dinner",
			"created_at": "2025-01-12T15:04:05Z"
		  }
		]`

		if err := os.WriteFile(filename, []byte(jsonTasks), 0o644); err != nil {
			t.Fatalf("Failed to write temp file: %v", err)
		}

		tasks, err := Load(filename)
		if err != nil {
			t.Fatalf("Failed to load tasks: %v", err)
		}

		if len(tasks) != 2 {
			t.Fatalf("Expected 2 tasks, got %d", len(tasks))
		}

		if tasks[0].ID != 1 {
			t.Errorf("Task[0].ID: got %d, want %d", tasks[0].ID, 1)
		}
		if tasks[0].Description != "Buy groceries" {
			t.Errorf("Task[0].Description: got %q, want %q", tasks[0].Description, "Buy groceries")
		}

		if tasks[1].ID != 2 {
			t.Errorf("Task[1].ID: got %d, want %d", tasks[1].ID, 2)
		}
		if tasks[1].Description != "Cook dinner" {
			t.Errorf("Task[1].Description: got %q, want %q", tasks[1].Description, "Cook dinner")
		}

	})

	t.Run("Invalid JSON returns error", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "invalid.json")

		invalidJSON := `[
		  {
			"id": 1,
			"description": "Broken JSON",
			"created_at": "2025-01-12T15:04:05Z"
		  }
		`
		if err := os.WriteFile(filename, []byte(invalidJSON), 0o644); err != nil {
			t.Fatalf("Failed to write temp file: %v", err)
		}

		tasks, err := Load(filename)
		if err == nil {
			t.Fatalf("Expected error for invalid JSON, got nil (tasks: %+v)", tasks)
		}
	})

	t.Run("Non-existing file returns empty slice", func(t *testing.T) {
		filename := "non-existing-file.json"

		tasks, err := Load(filename)
		if err != nil {
			t.Fatalf("Expected no error for non-existing file, got: %v", err)
		}

		if tasks == nil {
			t.Fatalf("Expected empty slice, got nil")
		}

		if len(tasks) != 0 {
			t.Fatalf("Expected 0 tasks, got %d", len(tasks))
		}
	})

	t.Run("Empty file returns empty slice", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "empty.json")

		tasks, err := Load(filename)
		if err != nil {
			t.Fatalf("Expected no error for empty file, got: %v", err)
		}

		if tasks == nil {
			t.Fatalf("Expected empty slice, got nil")
		}

		if len(tasks) != 0 {
			t.Fatalf("Expected 0 tasks, got %d", len(tasks))
		}
	})
}

func TestSave(t *testing.T) {
	t.Run("Valid tasks to empty file", func(t *testing.T) {
		dir := t.TempDir()
		filename := filepath.Join(dir, "tasks.json")

		t1 := time.Date(2025, 01, 12, 15, 4, 5, 0, time.UTC)
		tasks := []Task{
			{ID: 1, Description: "Buy groceries", CreatedAt: t1},
			{ID: 2, Description: "Cook dinner", CreatedAt: t1},
		}

		if err := Save(filename, tasks); err != nil {
			t.Fatalf("Save returned error: %v", err)
		}

		data, err := os.ReadFile(filename)
		if err != nil {
			t.Fatalf("failed to read saved file: %v", err)
		}

		var saved []Task
		if err := json.Unmarshal(data, &saved); err != nil {
			t.Fatalf("failed to unmarshal saved JSON: %v", err)
		}

		// Validate content
		if len(saved) != 2 {
			t.Fatalf("expected 2 tasks, got %d", len(saved))
		}

		if saved[0].ID != 1 || saved[0].Description != "Buy groceries" {
			t.Errorf("unexpected task[0]: %+v", saved[0])
		}
		if saved[1].ID != 2 || saved[1].Description != "Cook dinner" {
			t.Errorf("unexpected task[1]: %+v", saved[1])
		}

		if !saved[0].CreatedAt.Equal(t1) {
			t.Errorf("task[0].CreatedAt mismatch: got %v, want %v", saved[0].CreatedAt, t1)
		}
		if !saved[1].CreatedAt.Equal(t1) {
			t.Errorf("task[1].CreatedAt mismatch: got %v, want %v", saved[1].CreatedAt, t1)
		}
	})
}
