package functions

import (
	"fmt"

	ps "github.com/mitchellh/go-ps"
)

func showProcesses() ([]string, error) {
	res := []string{}
	result, err := ps.Processes()
	if err != nil {
		return nil, err
	}

	nameWidth := 25
	ppidWidth := 10
	pidWidth := 10

	for _, value := range result {
		res = append(res, fmt.Sprintf("Name: %-*s Ppid: %-*d Pid: %-*d", nameWidth, value.Executable(), ppidWidth, value.PPid(), pidWidth, value.Pid()))
	}
	return res, nil
}
