package functions

import (
	"errors"
	"os"
	"strconv"
)

func killByPid(line []string) error {
	if len(line) < 2 {
		return errors.New("not enough arguments")
	}
	pid, err := strconv.Atoi(line[1])
	if err != nil {
		return errors.New("invalid PID")
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return process.Signal(os.Kill)
}
