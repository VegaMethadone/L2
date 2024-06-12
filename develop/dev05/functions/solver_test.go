package functions

import (
	"bufio"
	"fmt"
	"log"
	"mangrep/structs"
	"os"
	"testing"
)

func readFileData() []string {
	lines := []string{}

	wd, err := os.Getwd()
	if err != nil {
		log.Printf("%v\n", err)
	}
	fmt.Println(wd)
	filename := "../test01.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("%v\n", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("%v\n", err)
	}

	return lines
}

func TestNumirateString(t *testing.T) {
	lines := readFileData()

	numirateString(&lines)

	for i, value := range lines {
		fmt.Printf("%d\t%s\n", i+1, value)
	}
}

func TestIgnoreCase(t *testing.T) {
	lines := readFileData()

	ignoreCase(&lines)

	for i, value := range lines {
		fmt.Printf("%d\t%s\n", i+1, value)
	}
}

func TestFindBefore(t *testing.T) {
	lines := readFileData()

	args := &structs.Args{
		After:     0,
		Before:    2,
		Context:   0,
		Count:     false,
		Ignore:    false,
		Invert:    false,
		Fixed:     true,
		Linenum:   false,
		Arguments: []string{"test"},
	}

	lines = findBefore(&lines, args)
	for i, value := range lines {
		fmt.Printf("%d\t%s\n", i+1, value)
	}
}

func TestFindAfter(t *testing.T) {
	lines := readFileData()

	args := &structs.Args{
		After:     2,
		Before:    0,
		Context:   0,
		Count:     false,
		Ignore:    false,
		Invert:    false,
		Fixed:     false,
		Linenum:   false,
		Arguments: []string{"test"},
	}

	lines = findAfter(&lines, args)
	for i, value := range lines {
		fmt.Printf("%d\t%s\n", i+1, value)
	}
}

func TestFindContext(t *testing.T) {
	lines := readFileData()

	args := &structs.Args{
		After:     0,
		Before:    0,
		Context:   2,
		Count:     false,
		Ignore:    false,
		Invert:    false,
		Fixed:     false,
		Linenum:   false,
		Arguments: []string{"test"},
	}

	lines = findContext(&lines, args)
	for i, value := range lines {
		fmt.Printf("%d\t%s\n", i+1, value)
	}
}

func TestCounter(t *testing.T) {
	lines := readFileData()

	args := &structs.Args{
		After:     0,
		Before:    0,
		Context:   0,
		Count:     true,
		Ignore:    false,
		Invert:    false,
		Fixed:     false,
		Linenum:   false,
		Arguments: []string{"test"},
	}
	ignoreCase(&lines)
	lines = countMatch(&lines, args)
	for i, value := range lines {
		fmt.Printf("%d\t%s\n", i+1, value)
	}
}
