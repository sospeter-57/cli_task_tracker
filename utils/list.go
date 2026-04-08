package utils

import (
	"fmt"
)

func ListAll() {
	if len(tasks) == 0 {
		fmt.Println("Uups! no tasks available. run: task-cli add <description>")
	} else {
		for _, task := range tasks {
			fmt.Println(task)
		}
	}
}

func ListDone() {
	if len(tasks) == 0 {
		fmt.Println("Uups! no tasks available. run: task-cli add <description>")
	} else {
		for _, task := range tasks {
			if task.TaskStatus == "done" {
				fmt.Println(task)
			}
		}
	}
}

func ListTodo() {
	if len(tasks) == 0 {
		fmt.Println("Uups! no tasks available. run: task-cli add <description>")
	} else {
		for _, task := range tasks {
			if task.TaskStatus == "todo" {
				fmt.Println(task)
			}
		}
	}
}

func ListInProgress() {
	if len(tasks) == 0 {
		fmt.Println("Uups! no tasks available. run: task-cli add <description>")
	} else {
		for _, task := range tasks {
			if task.TaskStatus == "in-progress" {
				fmt.Println(task)
			}
		}
	}
}	

func ListByStatus(status string) {
	switch status {
	case "done":
		ListDone()
	case "todo":
		ListTodo()
	case "in-progress":
		ListInProgress()
	default:
		fmt.Println("Invalid status provided. Valid status: done, todo, in-progress")
	}	
}