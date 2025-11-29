package handler

import (
	"github.com/dandimuzaki/project-app-task-list-cli-nama/cmd/dto"
	"github.com/dandimuzaki/project-app-task-list-cli-nama/service"
	"github.com/dandimuzaki/project-app-task-list-cli-nama/utils"
)

type TaskHandler interface {
	GetAllTask() ([]dto.TaskResponse, string)
	GetTaskByID(uint) (dto.TaskResponse, string)
	CreateTask(dto.CreateTaskRequest) (dto.MessageResponse)
	UpdateTask(uint, dto.UpdateTaskRequest) (dto.MessageResponse)
	DeleteTask(uint) (dto.MessageResponse)
}

type taskHandler struct {
	service service.TaskService
}

func NewTaskHandler(s service.TaskService) TaskHandler {
	return &taskHandler{
		service: s,
	}
}

func (h *taskHandler) CreateTask(req dto.CreateTaskRequest) (res dto.MessageResponse) {
	_, err := h.service.CreateTask(req)
	if err != nil {
		res.Message = utils.ErrorMessage(err)
	} else {
		res.Message = "\033[32m[SUCCESS] New task is created successfully\033[0m"
	}
	return res
}

func (h *taskHandler) GetAllTask() ([]dto.TaskResponse, string) {
	todoList, err := h.service.GetAllTask()
	if err != nil {
		return []dto.TaskResponse{}, utils.ErrorMessage(err)
	} else {
		var todoResponse []dto.TaskResponse
		for _, t := range todoList {
			task := *dto.NewTaskResponse(t.ID, t.Activity, t.Status, t.Priority)
			todoResponse = append(todoResponse, task)
		}
		return todoResponse, ""
	}
}

func (h *taskHandler) GetTaskByID(id uint) (dto.TaskResponse, string) {
	task, err := h.service.GetTaskByID(id)
	if err != nil {
		return dto.TaskResponse{}, utils.ErrorMessage(err)
	} else {
		res := dto.NewTaskResponse(task.ID, task.Activity, task.Status, task.Priority)
		return *res, ""
	}
}

func (h *taskHandler) UpdateTask(id uint, req dto.UpdateTaskRequest) (res dto.MessageResponse) {
	_, err := h.service.UpdateTask(id, req)
	if err != nil {
		res.Message = utils.ErrorMessage(err)
	} else {
		res.Message = "\033[32m[SUCCESS] Task is updated successfully\033[0m"
	}
	return res
}

func (h *taskHandler) DeleteTask(id uint) (res dto.MessageResponse) {
	err := h.service.DeleteTask(id)
	if err != nil {
		res.Message = utils.ErrorMessage(err)
	} else {
		res.Message = "\033[32m[SUCCESS] Task is deleted successfully\033[0m"
	}
	return res
}