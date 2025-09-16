package main

import (
	"encoding/json"
	"os"
)

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task

	err = json.Unmarshal(data, &tasks)

	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}
