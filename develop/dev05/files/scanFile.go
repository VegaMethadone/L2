package files

import (
	"bufio"
	"errors"
	"mangrep/structs"
	"os"
)

func GetLines(args *structs.Args) ([]string, error) {
	lines := []string{}

	if len(args.Arguments) == 1 {
		return nil, errors.New("file is not found in command line")
	} else {
		filename := args.Arguments[1]
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}
	return lines, nil
}
