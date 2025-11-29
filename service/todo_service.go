package service

import (
	"errors"
	"strings"
	"time"

	"github.com/dandimuzaki/project-app-task-list-cli-nama/cmd/dto"
	"github.com/dandimuzaki/project-app-task-list-cli-nama/model"
	"github.com/dandimuzaki/project-app-task-list-cli-nama/utils"
)

type TaskService interface {
	CreateTask(dto.CreateTaskRequest) (model.Task, error)
	GetAllTask(dto.QueryRequest) ([]model.Task, error)
	GetTaskByID(uint) (*model.Task, error)
	UpdateTask(uint, dto.UpdateTaskRequest) (model.Task, error)
	DeleteTask(id uint) error
}

type taskService struct{}

func NewTaskService() TaskService {
	return &taskService{}
}

func (s *taskService) CreateTask(req dto.CreateTaskRequest) (model.Task, error) {
	// Validate content of request
	if strings.TrimSpace(req.Activity) == "" ||
		strings.TrimSpace(req.Priority) == "" {
		return model.Task{}, errors.New("activity and priority are required")
	}

	// Read data from existing json file
	todoList, err := utils.ReadTodoFromFile()
	if err != nil {
		return model.Task{}, err
	}

	// Generate simple ID (max ID + 1)
	newID := uint(1)
	for _, t := range todoList {
		if uint(t.ID) >= newID {
			newID = t.ID + 1
		}
	}
	
	newTask := model.Task{
		ID: newID,
		Activity: utils.Capitalize(req.Activity),
		Status: "On Progress",
		Priority: utils.Uppercase(req.Priority),
		CreatedAt: time.Now(),
	}

	todoList = append(todoList, newTask)

	if err := utils.WriteTodoToFile(todoList); err != nil {
		return model.Task{}, err
	}

	return newTask, nil
}


func (s *taskService) GetAllTask(q dto.QueryRequest) ([]model.Task, error) {
	todoList, err := utils.ReadTodoFromFile()
	if err != nil {
		return []model.Task{}, err
	}
	var filtered []model.Task
	if q.Keyword != "" {
		for _, t := range todoList {
			if strings.Contains(strings.ToLower(t.Activity), strings.ToLower(q.Keyword)) {
				filtered = append(filtered, t)
			}
		}
		return filtered, nil
	}
	if q.Status != "" {
		for _, t := range todoList {
			if strings.Contains(strings.ToLower(t.Status), strings.ToLower(q.Status)) {
				filtered = append(filtered, t)
			}
		}
		return filtered, nil
	}
	if q.Priority != "" {
		for _, t := range todoList {
			if strings.Contains(strings.ToLower(t.Priority), strings.ToLower(q.Priority)) {
				filtered = append(filtered, t)
			}
		}
		return filtered, nil
	}
	return todoList, nil
}

func (s *taskService) GetTaskByID(id uint) (*model.Task, error) {
	todoList, err := utils.ReadTodoFromFile()
	if err != nil {
		return &model.Task{}, err
	}

	for _, t := range todoList {
		if t.ID == id {
			copy := t
			return &copy, nil
		}
	}

	return nil, utils.ErrNotFound
}

func (s *taskService) UpdateTask(id uint, req dto.UpdateTaskRequest) (model.Task, error) {
	todoList, err := utils.ReadTodoFromFile()
	if err != nil {
		return model.Task{}, err
	}

	found := false
	var task model.Task
	for _, t := range todoList {
		if t.ID == id {
			found = true
			task = t
		}
	}
	if !found {
		return model.Task{}, utils.ErrNotFound
	}

	if task != (model.Task{}) {
		if req.Activity != "" {
			task.Activity = utils.Capitalize(req.Activity)
		}
		if req.Priority != "" {
			task.Priority = utils.Uppercase(req.Priority)
		}
		if req.Status != "" {
			task.Status = utils.Uppercase(req.Status)
		}
	}

	todoList[id-1] = task

	return task, utils.WriteTodoToFile(todoList)
}

func (s *taskService) DeleteTask(id uint) error {
	todoList, err := utils.ReadTodoFromFile()
	if err != nil {
		return utils.ErrFile
	}

	newList := make([]model.Task, 0, len(todoList))
	found := false
	for _, task := range todoList {
		if task.ID == id {
			found = true
			continue
		}
		newList = append(newList, task)
	}

	if !found {
		return utils.ErrNotFound
	}

	return utils.WriteTodoToFile(newList)
}