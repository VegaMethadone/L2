package functions

import (
	"errors"
	"os"
	"runtime"
	"strings"
)

func cd(line []string) error {
	if len(line) < 2 {
		return errors.New("not enough arguments")
	}

	wg, err := os.Getwd()
	if err != nil {
		return err
	}

	separator := string(os.PathSeparator)
	if runtime.GOOS == "windows" {
		separator = "\\"
	}

	splitedWd := strings.Split(wg, separator)

	wantPath := strings.Split(line[1], "/")
	if runtime.GOOS == "windows" {
		wantPath = strings.Split(line[1], "\\")
	}

	for _, part := range wantPath {
		if part == ".." {
			if len(splitedWd) > 0 {
				splitedWd = splitedWd[:len(splitedWd)-1]
			}
		} else {
			splitedWd = append(splitedWd, part)
		}
	}

	newPath := strings.Join(splitedWd, separator)

	return os.Chdir(newPath)
}
