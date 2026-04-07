package main

import (
		"os"
		"fmt"
		"log"
		"strconv"
		"cli_task_tracker/parser"
		"cli_task_tracker/utils"
)

func main() {

	err := utils.InitializeTasks()
	if err != nil {
		log.Fatalln(err)
	}
	
		args := os.Args[1:]
		command, additional, err := parser.ParseArguments(args);
		if err != nil {
			log.Fatalln(err)
	}
		switch command {
		case "add":
			err := utils.AddTask(additional[0])
			if err != nil {
				fmt.Println(err)
			}
		case "list":
			utils.ListTasks()
		case "delete":
			id, err := strconv.ParseInt(additional[0], 10, 64)
			if err != nil {
				log.Fatalln("Failed to parse task id number")
			}
			err = utils.DeleteTask(uint32(id))
			if err != nil {
				fmt.Println(err)
			}
		case "update":
			id, err := strconv.ParseInt(additional[0], 10, 32)
			if err != nil {
				log.Fatalln("Failed to parse task id number")
			}
			err = utils.UpdateTask(uint32(id), additional[1])
			if err != nil {
				fmt.Println(err)
			}
		case "mark-done":
			id, err := strconv.ParseInt(additional[0], 10, 64)
			if err != nil {
				log.Fatalln("Failed to parse task id number")
			}
			err = utils.MarkDone(uint32(id))
			fmt.Println(err)
		case "mark-in-progress":
			id, err := strconv.ParseInt(additional[0], 10, 64)
			if err != nil {
				log.Fatalln("Failed to parse task id number")
			}
			err = utils.MarkInProgress(uint32(id))
			if err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("the fucking default case")
		}
}

