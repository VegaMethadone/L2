package main

import (
	"bufio"
	"fmt"
	"mancut/arguments"
	"mancut/functions"
	"os"
	"strings"
)

func main() {
	args, err := arguments.GetArguments()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fields, err := functions.ParseFields(args.Fields)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	line, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	line = strings.TrimSpace(line)

	if args.Separated && !strings.Contains(line, args.Delimiter) {
		return
	}

	columns := strings.Split(line, args.Delimiter)
	selectedColumns := functions.SelectColumn(columns, fields)

	fmt.Println(strings.Join(selectedColumns, args.Delimiter))
}
