package utils

import (
	// "encoding/json"
	"time"
	// "os"
	// "errors"

	"cli_task_tracker/models"
)

func AddTask(description string) error {
	var err error

	var taskId uint32 = 0
	currentTime := time.Now().Format("2006-01-02 07:00:00")
	if !(len(tasks) == 0) {
		taskId = tasks[len(tasks)-1].Id + 1
	}
	task := models.Task{
		Id:          taskId,
		Description: description,
		TaskStatus:  models.StatusTodo,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	tasks = append(tasks, task)
	if err = UpdateLogFile(); err != nil {
		return err
	}
	return nil
}
