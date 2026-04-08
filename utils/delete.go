package utils

import (
	"errors"
)

func DeleteTask(taskId uint32) error {

	for taskIndex, task := range tasks {
		if task.Id == taskId {
			firstHalf := tasks[:taskIndex]
			lastHalf := tasks[taskIndex+1:]
			tasks = firstHalf
			tasks = append(tasks, lastHalf...)

			if err := UpdateLogFile(); err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("A task with that Id wasn't found")
}
