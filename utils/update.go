package utils

import (
	"errors"

	"cli_task_tracker/models"
)

func UpdateTask(taskId uint32, description string) error {

	var err error

	if len(description) == 0 {
		err = errors.New("Sorry, description can't be nil.")
		return err
	} else {
		for _, task := range tasks {
			if task.Id == taskId {
				task.Description = description
				
				err = UpdateLogFile()
				if err != nil {
					return err
				}
				return nil
			}
		}
		err = errors.New("Sorry, a task with that id wasn't found")
		return err
	}
}


func MarkInProgress(taskId uint32) error {

	for _, task := range tasks {
		if task.Id == taskId {
			task.TaskStatus = models.StatusInProgress
			return nil
		}
	}

	return errors.New("Sorry, a task with that task id doesn't exist")
}

func MarkDone(taskId uint32) error {

	for _, task := range tasks {
		if task.Id == taskId {
			task.TaskStatus = models.Done
			return nil
		}
	}

	return errors.New("Sorry, as task with that task id doesn't exist")
}