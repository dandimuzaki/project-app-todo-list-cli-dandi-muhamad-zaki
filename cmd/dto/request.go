package dto

type CreateTaskRequest struct {
	Activity, Priority string
}

func NewCreateTaskRequest(activity, priority string) *CreateTaskRequest {
	return &CreateTaskRequest{activity, priority}
}

type UpdateTaskRequest struct {
	Activity, Status, Priority string
}

func NewUpdateTaskRequest(activity, status, priority string) *UpdateTaskRequest {
	return &UpdateTaskRequest{activity, status, priority}
}