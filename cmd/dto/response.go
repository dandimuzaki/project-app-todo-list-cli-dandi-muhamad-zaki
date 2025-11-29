package dto

type TaskResponse struct {
	ID                         uint
	Activity, Status, Priority string
}

func NewTaskResponse(id uint, activity, status, priority string) *TaskResponse {
	return &TaskResponse{id, activity, status, priority}
}

type MessageResponse struct {
	Message string
}