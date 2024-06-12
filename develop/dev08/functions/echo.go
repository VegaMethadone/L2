package functions

import (
	"errors"
	"fmt"
	"strings"
)

func echo(line []string) error {
	if len(line) < 2 {
		return errors.New("not enough arguments")
	}
	fmt.Print(strings.Join(line[1:], " "))
	return nil
}
