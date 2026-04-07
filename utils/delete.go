package utils

import (
	"errors"
)
func DeleteTask(taskId uint32) error {
	
	for taskIndex, task := range tasks {
		if task.Id == taskId {
			firstHalf := tasks[:taskIndex-1]
			lastHalf := tasks[taskIndex:]
			tasks = firstHalf
			tasks = append(tasks, lastHalf...)
			return nil
		}
	}

	return errors.New("A task with that Id wasn't found")
}