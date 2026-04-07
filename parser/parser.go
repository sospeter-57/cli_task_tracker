package parser

import (
//		"strings"
		"errors"
)

func ParseArguments(args []string) (string, []string, error) {
		additionalArguments := []string{}

		var err error

		argsCount := len(args)
		if argsCount < 1 {
				err = errors.New("No arguments supplied")
			return "",[]string{}, err
	}

	tasks := []string{"add", "update", "delete", "list", "mark-in-progress", "mark-done"}

	command := args[0]

	for _, task := range tasks {
			if task == command {
				additionalArguments = args[1:]
					return command,additionalArguments, nil
			}
	}
	err = errors.New("Invalid command parsed.\nValid commands: add, update, delete, list")

	return "",[]string{}, err
}


