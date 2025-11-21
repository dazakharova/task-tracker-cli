package tasks

import (
	"encoding/json"
	"os"
)

func Load(file string) ([]Task, error) {
	var tasks []Task

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []Task{}, nil
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func Save(file string, tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0644)
}
