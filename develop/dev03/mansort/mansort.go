package mansort

import (
	"mansort/arguments"
	"mansort/files"
	"mansort/functions"
	"strings"
)

func Sort() (string, error) {
	args, err := arguments.GetFlags()
	if err != nil {
		return "", err
	}

	lines, err := files.GetLines(args)
	if err != nil {
		return "", err
	}

	res := functions.Solver(lines, args)

	return strings.Join(res, "\n"), nil
}
