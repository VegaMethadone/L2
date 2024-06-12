package files

import (
	"bufio"
	"errors"
	"mansort/structs"
	"os"
)

func GetLines(args *structs.Args) ([]string, error) {
	lines := []string{}

	if len(args.Files) == 0 {
		return nil, errors.New("file is not found in command line")
	} else {
		for _, filename := range args.Files {
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
	}
	return lines, nil
}
