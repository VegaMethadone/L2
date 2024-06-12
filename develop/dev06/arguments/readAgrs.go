package arguments

import (
	"errors"
	"flag"
	"mancut/structs"
)

func GetArguments() (*structs.Args, error) {
	columns := flag.String("f", "", "select columns")
	delimiter := flag.String("d", "\t", "use different delimiter")
	separate := flag.Bool("s", false, "only lines with delimiter")

	flag.Parse()

	args := &structs.Args{
		Fields:    *columns,
		Delimiter: *delimiter,
		Separated: *separate,
	}

	if args.Fields == "" {
		return nil, errors.New("flag -f is required")
	}

	args.Command = flag.Args()

	return args, nil
}
