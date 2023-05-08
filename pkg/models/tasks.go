package models

type Task struct {
	Id          int    `json:"id"`
	CreateTime  string `json:"create_time"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Level       int    `json:"level"`
	// Solved      bool   `json:"solved"`
}

type AddTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}

type SolveTaskRequest struct {
	TaskId int `json:"task_id"`
	UserId int `json:"user_id"`
}
