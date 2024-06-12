package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"testing"
)

type testCase struct {
	args    []string
	command string
}

func TestManCut(t *testing.T) {
	cases := []testCase{
		{[]string{"-f", "1,3"}, "col1\tcol2\tcol3\tcol4\tcol5\n"},
		{[]string{"-d", ",", "-f", "2"}, "col1,col2,col3\n"},
		{[]string{"-d", " ", "-f", "3-5"}, "col1 col2 col3 col4 col5 col6 col7\n"},
		{[]string{"-d", ",", "-f", "1", "-s"}, "col1,col2\n"},
	}

	for _, tc := range cases {
		cmd := exec.Command("go", append([]string{"run", "main.go"}, tc.args...)...)

		stdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatalf("Failed to create stdin pipe: %v", err)
		}

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Start(); err != nil {
			t.Fatalf("Failed to start command: %v", err)
		}

		_, err = io.WriteString(stdin, tc.command)
		if err != nil {
			t.Fatalf("Failed to write to stdin: %v", err)
		}
		stdin.Close()

		if err := cmd.Wait(); err != nil {
			t.Fatalf("Command failed: %v, stderr: %s", err, stderr.String())
		}

		output := stdout.String()

		fullCommand := fmt.Sprintf("go run main.go %s", strings.Join(tc.args, " "))
		fmt.Println("------------------------------------------")
		fmt.Printf("Command: %s\n", fullCommand)
		fmt.Printf("Input: %s", tc.command)
		fmt.Printf("Output: %s\n", output)
	}
}
