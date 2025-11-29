package model

import "time"

type Task struct {
	ID uint `json:"id"`
	Activity string `json:"activity"`
	Status string `json:"status"`
	Priority string `json:"priority"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}