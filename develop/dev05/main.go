package main

import (
	"fmt"
	"mangrep/grep"
	"os"
)

func main() {
	result, err := grep.Grep()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, str := range result {
		fmt.Println(str)
	}
}
