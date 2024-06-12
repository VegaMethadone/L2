package main

import (
	"fmt"
	"mansort/mansort"
	"os"
)

func main() {
	result, err := mansort.Sort()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(result)
}
