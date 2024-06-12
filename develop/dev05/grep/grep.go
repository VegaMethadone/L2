package grep

import (
	"mangrep/arguments"
	"mangrep/files"
	"mangrep/functions"
)

func Grep() ([]string, error) {
	args, err := arguments.GetArgs()
	if err != nil {
		return nil, err
	}
	lines, err := files.GetLines(args)
	if err != nil {
		return nil, err
	}

	return functions.TheGod(&lines, args), nil
}
