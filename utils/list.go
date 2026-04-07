package utils

import (
	"fmt"
)

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("Uups! no tasks available. run: task-cli add <description>")
	} else {
		for index, task := range tasks {
			fmt.Println(index, task)
		}
	}
}