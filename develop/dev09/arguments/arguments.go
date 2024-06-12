package arguments

import (
	"errors"
	"flag"
	"manwget/structs"
)

func GetArgs() (*structs.Args, error) {
	name := flag.String("o", "", "name of file")
	recursion := flag.Int("r", 0, "deapth  of recursion")

	flag.Parse()

	args := &structs.Args{
		Name:      *name,
		Recursion: *recursion,
	}

	args.Url = flag.Args()

	if args.Recursion < 0 {
		return nil, errors.New("invalid -r value: less than 0")
	}

	return args, nil
}
