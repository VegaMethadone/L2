package main

import (
	"bufio"
	"fmt"
	"os"
	"pog/parser"
)

func main() {
	for {

		fmt.Printf("$: ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		_, err = parser.ReadLine(line)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
