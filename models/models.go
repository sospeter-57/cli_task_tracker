package models

type Task struct {
	Id          uint32 `json:"id"`
	Description string `json:"description"`
	TaskStatus  Status `json:"task_status"`
	CreatedAt   string `json:"crated_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	Done             Status = "done"
)
