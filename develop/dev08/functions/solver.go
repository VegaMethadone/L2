package functions

import (
	"fmt"
	"pog/structs"
)

func Solver(line []string) error {
	switch {
	case line[0] == structs.CD:
		return cd(line)
	case line[0] == structs.PWD:
		dir, err := currentDir()
		if err != nil {
			return err
		}
		fmt.Println(dir)
	case line[0] == structs.ECHO:
		return echo(line)
	case line[0] == structs.KILL:
		return killByPid(line)
	case line[0] == structs.PS:
		lines, err := showProcesses()
		if err != nil {
			return err
		}
		for _, l := range lines {
			fmt.Println(l)
		}
		return nil
	}

	return nil
}
