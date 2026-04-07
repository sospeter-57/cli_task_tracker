package utils

import (
	"os"
	"io"
	"encoding/json"
	"cli_task_tracker/models"
)

var tasks []models.Task 
const LogFile string = "logs.json"

func InitializeTasks() error {

	file, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	return nil
}

func UpdateLogFile() error {
	file, err := os.OpenFile(LogFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
		err = encoder.Encode(tasks)

		return nil
}