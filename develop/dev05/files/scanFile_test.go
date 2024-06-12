package files

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestGetLines(t *testing.T) {
	lines := []string{}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	fmt.Println(wd)
	filename := "../test01.txt"
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("%v\n", err)
	}

	for i, value := range lines {
		fmt.Printf("%d\t%s\n", i+1, value)
	}
}
