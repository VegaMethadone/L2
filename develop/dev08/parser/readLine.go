package parser

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"pog/functions"
	"pog/structs"
	"strings"
)

func ReadLine(line string) ([]string, error) {

	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return nil, errors.New("not enough arguments")
	}

	commands := strings.Split(line, "|")
	if len(commands) > 1 {
		for _, value := range commands {
			if _, err := ReadLine(value); err != nil {
				return nil, err
			}
		}
		return nil, nil
	}

	tmp := strings.Fields(line)
	if len(tmp) == 0 {
		return nil, errors.New("not enough arguments")
	}

	command := tmp[0]
	args := tmp[1:]

	switch command {
	case structs.CD, structs.PWD, structs.ECHO, structs.KILL, structs.PS:
		err := functions.Solver(append([]string{command}, args...))
		if err != nil {
			return nil, err
		}
	case "exit", "quit":
		os.Exit(0)
	default:
		cmd := exec.Command(command, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
			return nil, err
		}
	}

	return append([]string{command}, args...), nil
}
