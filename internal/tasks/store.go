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
