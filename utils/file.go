package utils

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/dandimuzaki/project-app-task-list-cli-nama/model"
)

const todoFilePath = "data/todoList.json"

// cek directory and file
func EnsureTodoFile() error {
	_, err := os.Stat(todoFilePath)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll("data", 0755); err != nil {
			return err
		}
		return os.WriteFile(todoFilePath, []byte("[]"), 0644)
	}
	return nil
}

// read file
func ReadTodoFromFile() ([]model.Task, error) {
	if err := EnsureTodoFile(); err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(todoFilePath)
	if err != nil {
		return nil, err
	}

	var todoList []model.Task
	if err := json.Unmarshal(bytes, &todoList); err != nil {
		return nil, err
	}

	return todoList, nil
}

// write file
func WriteTodoToFile(todoList []model.Task) error {
	bytes, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(todoFilePath, bytes, 0644)
}
