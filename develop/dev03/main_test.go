package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

type testCase struct {
	a []string
	f []string
}

func TestSort(t *testing.T) {
	cases := []testCase{
		{a: []string{"-r"}, f: []string{"test01.txt"}},
		{a: []string{}, f: []string{"test01.txt"}},
		{a: []string{"-r"}, f: []string{"test06.txt"}},
		{a: []string{}, f: []string{"test06.txt"}},
		{a: []string{"-k", "2", "-n"}, f: []string{"test02.txt"}},
		{a: []string{"-k", "2", "-r", "-n"}, f: []string{"test02.txt"}},
		{a: []string{"-k", "3", "-M"}, f: []string{"test03.txt"}},
		{a: []string{"-k", "3", "-M", "-r"}, f: []string{"test03.txt"}},
		{a: []string{"-u", "-k", "3", "-M"}, f: []string{"test04.txt"}},
		{a: []string{"-u", "-k", "2", "-r", "-n"}, f: []string{"test05.txt"}},
	}

	for _, tc := range cases {
		commands := []string{"run", "main.go"}
		commands = append(commands, tc.a...)
		commands = append(commands, tc.f...)

		cmd, err := exec.Command("go", commands...).CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not exec command: %v", err)
			os.Exit(1)
		}
		fmt.Println("----------------------------------------------------")
		fmt.Println(strings.Join(append([]string{"go"}, commands...), " "))
		fmt.Println("----------------------------------------------------")
		fmt.Println(string(cmd))

	}
}
