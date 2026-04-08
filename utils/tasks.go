package utils

import (
	"cli_task_tracker/models"
	"encoding/json"
	"fmt"
	"io"
	"os"
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

	file, err := os.OpenFile(LogFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Sorry, and error occurred while opening the log file: %v", err)
	}
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		file.Close()
		return fmt.Errorf("Sorry, an error occurred while updating the log file: %v", err)
	}
	return nil

}
