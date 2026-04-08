package utils

import (
	"cli_task_tracker/models"
	"errors"
)

func UpdateTask(taskId uint32, description string) error {

	var err error

	if len(description) == 0 {
		err = errors.New("Sorry, description can't be nil.")
		return err
	} else {
		for index := range tasks {
			if tasks[index].Id == taskId {
				tasks[index].Description = description

				if err = UpdateLogFile(); err != nil {
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

	for index := range tasks {
		if tasks[index].Id == taskId {
			tasks[index].TaskStatus = models.StatusInProgress
			if err := UpdateLogFile(); err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("Sorry, a task with that task id doesn't exist")
}

func MarkDone(taskId uint32) error {

	for index := range tasks {
		if tasks[index].Id == taskId {
			tasks[index].TaskStatus = models.Done
			if err := UpdateLogFile(); err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("Sorry, as task with that task id doesn't exist")
}
