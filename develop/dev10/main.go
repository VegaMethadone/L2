package main

import (
	"fmt"
	"os"
	"telnet/arguments"
	"telnet/functions"
)

func main() {
	args, err := arguments.GetArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = functions.Telnet(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
