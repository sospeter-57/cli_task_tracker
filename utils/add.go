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

	// file, err := os.OpenFile(LogFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	// encoder := json.NewEncoder(file)

	var taskId uint32 = 0
	currentTime := time.Now().Format("2006-01-02T15:04:05 -07:00:00")
	if !(len(tasks) == 0) {
		taskId = tasks[len(tasks)-1].Id + 1
	}
	task := models.Task {
		Id: taskId,
		Description: description,
		TaskStatus: models.StatusTodo,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	tasks = append(tasks, task)
	err = UpdateLogFile()

	// err = encoder.Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}

